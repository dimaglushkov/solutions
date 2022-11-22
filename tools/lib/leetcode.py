import os

import leetcode
import pandas as pd

from . import shared

CHARTS = ['by_difficulty', 'by_tags']
META_FILE = ".meta.csv"
LANG_SPECS = shared.get_lang_specs()


def _get_title_slugs(problems: list, lang: str, d: str) -> dict:
    slugs = dict()
    for problem in problems:
        if problem.startswith('http'):
            slug = problem
            if problem.endswith('/'):
                slug = slug[:-1].split('/')[-1]
            slugs[slug] = os.path.join(d, slug, f'{slug}.{LANG_SPECS[lang]["ext"]}')
            if not os.path.exists(os.path.join(d, slug)):
                os.mkdir(os.path.join(d, slug))
        else:
            slugs[problem] = os.path.join(d, problem, f'{problem}.{LANG_SPECS[lang]["ext"]}')
            if not os.path.exists(os.path.join(d, problem)):
                os.mkdir(os.path.join(d, problem))

    existing_solutions = list()
    for k, v in slugs.items():
        if os.path.exists(v):
            print(f'Skipping {k}, {v} already exists.')
            existing_solutions.append(k)
    for i in existing_solutions:
        del slugs[i]

    return slugs


def _get_slug_data(name: str) -> dict:
    configuration = leetcode.Configuration()
    configuration.api_key["Referer"] = "https://leetcode.com"
    configuration.debug = False
    api_instance = leetcode.DefaultApi(leetcode.ApiClient(configuration))
    graphql_request = leetcode.GraphqlQuery(
        query="""
                query getQuestionDetail($titleSlug: String!) {
                  question(titleSlug: $titleSlug) {
                    questionId
                    questionFrontendId
                    boundTopicId
                    title
                    content
                    translatedTitle
                    difficulty
                    similarQuestions
                    topicTags {
                      name
                      slug
                    }
                    codeSnippets {
                      lang
                      langSlug
                      code
                      __typename
                    }
                    stats
                    codeDefinition
                    hints
                    solution {
                      id
                      canSeeDetail
                      __typename
                    }
                    status
                    sampleTestCase
                    enableRunCode
                    metaData
                    translatedContent
                    judgerAvailable
                    judgeType
                    mysqlSchemas
                    enableTestMode
                    envInfo
                    __typename
                  }
                }
            """,
        variables=leetcode.GraphqlQueryGetQuestionDetailVariables(title_slug=name),
        operation_name="getQuestionDetail",
    )
    return api_instance.graphql_post(body=graphql_request).to_dict()


def _get_test_cases(data: dict) -> list:
    tests = list()
    translation_table = {
        '<strong>': '',
        '</strong>': '',
        '&quot;': '"',
        "Input: ": '',
        "Output: ": '',
        '<pre>': ''
    }
    desc = data['content']
    rows = desc.split('\n')
    inputs_sets = [d for d in rows if 'Input:' in d]
    outputs = [d for d in rows if 'Output:' in d]

    if len(inputs_sets) != len(outputs):
        print(f"Warning: different number of input and output examples for task")
        return tests

    for ttk, ttv in translation_table.items():
        for i, s in enumerate(inputs_sets):
            inputs_sets[i] = s.replace(ttk, ttv)
        for i, s in enumerate(outputs):
            outputs[i] = s.replace(ttk, ttv)

    for j, inps in enumerate(inputs_sets):
        variables = {i.split(' = ')[0]: i.split(' = ')[1] for i in inps.split(', ')}
        tests.append({
            'input': variables,
            'output': outputs[j]
        })
    return tests


def _generate_golang_test_code(tests: list, code: str) -> str:
    tester = """\n
func main() {
    testCases := []struct {
%__inputs__%
%__output__%
    }{
%__testcases__%\t}

%__runner__%
}
"""
    if '/**' in code:
        code = code.split('*/\n')[-1]

    func_decl = code.split('\n')[0]
    func_name = func_decl.replace('func ', '').split('(')[0]
    func_params = {p.split(' ')[0]: ''.join(p.split(' ')[1:]) for p in
                   func_decl.split('(')[1].split(')')[0].split(', ')}
    func_return = func_decl.split(')')[-1].replace('{', '').strip()

    inputs = list()
    for n, t in func_params.items():
        inputs.append(f"\t\t{n} {t}")
    output = f"\t\twant {func_return}"

    testcases = list()
    for t in tests:
        tc = list()
        for vn, vt in func_params.items():
            tc.append(f'\t\t\t{vn}: {vt if "[]" in vt else ""} {t["input"][vn].replace("[", "{").replace("]", "}")},\n')
        tc.append(
            f'\t\t\twant: {func_return if "[]" in func_return else ""} '
            f'{t["output"].replace("[", "{").replace("]", "}")},\n'
        )
        tc_str = "".join(tc)
        testcases.append("\t\t{\n" + tc_str + "\t\t},\n")

    func_call = f"{func_name}({', '.join(['tc.' + p for p in func_params.keys()])})"
    runner = """
    successes := 0
    for _, tc := range testCases {
        x := %__func_call__%
        status := "ERROR"
        if fmt.Sprint(x) == fmt.Sprint(tc.want) {
            status = "OK"
            successes++
        }
        fmt.Println(status, "\tExpected: ", tc.want, " Actual: ", x)
    }
    if l := len(testCases); successes == len(testCases) {
        fmt.Printf("===\\nSUCCESS: %d of %d tests ended successfully\\n", successes, l)
    } else {
        fmt.Printf("===\\nFAIL: %d tests failed\\n", l - successes)
    }
""".replace("%__func_call__%", func_call)

    tester = tester \
        .replace("%__inputs__%", "\n".join(inputs)) \
        .replace("%__output__%", output) \
        .replace("%__testcases__%", "".join(testcases)) \
        .replace("%__runner__%", runner)
    return tester


def _create_code_template(slug: str, file: str, lang: str, data: dict) -> None:
    source_link = f'https://leetcode.com/problems/{slug}/'
    code_snippet = ''

    for snippet in data['code_snippets']:
        if snippet['lang_slug'] == lang:
            code_snippet = snippet['code']
            break

    code = 'package main\n\nimport (\n\t"fmt"\n)\n'
    if 'prefix' in LANG_SPECS[lang].keys():
        code = LANG_SPECS[lang]['prefix']

    code += f'{LANG_SPECS[lang]["com"]} source: {source_link}\n\n'
    code += code_snippet

    # experimental feature, tested only with golang
    if lang == 'golang':
        try:
            code += _generate_golang_test_code(_get_test_cases(data), code_snippet)
        except Exception as e:
            print("Was not able to generate test code: " + slug)

    with open(file, 'w') as code_file:
        code_file.write(code)


def _clear_leetcode_meta_file(data: dict, sol_dir: str):
    # keeping meta.csv actual by removing problems with no solution
    # obviously not the prettiest way to do it
    solutions = dict()
    for pkg in os.listdir(sol_dir):
        pkg_path = os.path.join(sol_dir, pkg)
        if os.path.isdir(pkg_path):
            for f in os.listdir(pkg_path):
                problem = f.split('.')[0]
                if problem not in solutions:
                    solutions[problem] = [f.split('.')[1]]
                else:
                    solutions[problem].append(f.split('.')[1])

    problem_to_del = list()
    lang_to_del = dict()

    for problem, meta in data.items():
        if problem not in solutions:
            problem_to_del.append(problem)
            continue

        for lang_id, old_lang in enumerate(meta['lang']):
            if LANG_SPECS[old_lang]['ext'] not in solutions[problem]:
                if problem not in lang_to_del:
                    lang_to_del[problem] = [lang_id]
                else:
                    lang_to_del[problem].append(lang_id)

    for problem, langs in lang_to_del.items():
        for lang_id in langs:
            del data[problem]['lang'][lang_id]

    for problem in problem_to_del:
        del data[problem]

    pd.DataFrame.from_dict(data).transpose().to_csv(os.path.join(sol_dir, META_FILE), index_label='slug', index=True)


def _generate_leetcode_readme(data: dict, sol_dir: str):
    header = '## leetcode'
    problems_table = '|Problem|Solution|Difficulty|Tags|\n|-|-|-|-|\n'

    # Getting simple stats and generating rows for the table
    for problem, meta in data.items():
        solutions_links = ', '.join(
            [f'[{lang}](/leetcode/{problem}/{problem}.{LANG_SPECS[lang]["ext"]})' for lang in meta['lang']])
        problems_table += \
            f'| [{meta["id"]}. {problem.replace("-", " ").capitalize()}](https://leetcode.com/problems/{problem}/) ' \
            f'| {solutions_links} ' \
            f'| {meta["difficulty"]} ' \
            f'| {", ".join(meta["tags"])} |\n'
    stats_str = shared.generate_svg_stats(data, sol_dir, CHARTS)
    readme = header + stats_str + problems_table
    with open(os.path.join(sol_dir, 'README.md'), 'w') as readme_file:
        readme_file.write(readme)


def _update_meta_file(slug: str, lang: str, slug_data: dict, sol_dir: str) -> None:
    meta_file = os.path.join(sol_dir, META_FILE)
    if not os.path.exists(meta_file):
        with open(meta_file, 'w') as f:
            f.write('slug,id,difficulty,tags,lang\n')
    data = pd.read_csv(meta_file, index_col='slug', converters={'lang': pd.eval}).to_dict('index')

    # adding new problem or solution
    if slug not in data.keys():
        data[slug] = {
            'id': slug_data['question_frontend_id'],
            'difficulty': slug_data['difficulty'],
            'tags': [tag['name'] for tag in slug_data['topic_tags']],
            'lang': [lang]
        }
    else:
        data[slug]['lang'].append(lang)

    pd.DataFrame.from_dict(data).transpose().to_csv(meta_file, index_label='slug', index=True)


def pull(problems: list, lang: str, sol_dir: str):
    if lang not in LANG_SPECS:
        print(f"Language {lang} is not supported yet\n"
              "In order to add it you need to add its extension to tools/lib/cf-problem-puller.py, "
              "look for LANG_SPECIFICS "
              "& create corresponding template file in tools/templates.")
        return

    if problems != ['.']:
        slugs = _get_title_slugs(problems, lang, sol_dir)
        if len(slugs) == 0:
            print("Nothing to do")
            return

        for slug, file in slugs.items():
            slug_data = _get_slug_data(slug)['data']['question']
            _create_code_template(slug, file, lang, slug_data)
            _update_meta_file(slug, lang, slug_data, sol_dir)

    data = pd.read_csv(
        os.path.join(sol_dir, META_FILE), index_col='slug', converters={'lang': pd.eval, 'tags': pd.eval}
    ).to_dict('index')
    _clear_leetcode_meta_file(data, sol_dir)
    _generate_leetcode_readme(data, sol_dir)


def delete(problems: list, lang: str, sol_dir: str):
    for problem in problems:
        if problem.startswith('http'):
            if not problem.endswith('/'):
                problem += '/'
            problem = problem[:-1].split('/')[-1]

        os.remove(os.path.join(sol_dir, problem, f'{problem}.{LANG_SPECS[lang]["ext"]}'))
        os.rmdir(os.path.join(sol_dir, problem))

    data = pd.read_csv(
        os.path.join(sol_dir, META_FILE), index_col='slug', converters={'lang': pd.eval, 'tags': pd.eval}
    ).to_dict('index')
    _clear_leetcode_meta_file(data, sol_dir)
    _generate_leetcode_readme(data, sol_dir)
