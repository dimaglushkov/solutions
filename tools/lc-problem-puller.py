import argparse
import leetcode
import os
import pandas as pd

lang_specifics = {
    'golang': {
        'ext': 'go',
        'com': '//',
        'prefix': 'package main\n\nimport "fmt"\n\n',
    },
    'python': {
        'ext': 'py',
        'com': '#'
    },
    'bash': {
        'ext': 'sh',
        'com': '#'
    }
}


def get_title_slugs(problems: list, lang: str, d: str) -> dict:
    slugs = dict()
    for problem in problems:
        if problem.startswith('http'):
            slug = problem
            if problem.endswith('/'):
                slug = slug[:-1].split('/')[-1]
            slugs[slug] = os.path.join(d, f'{slug}.{lang_specifics[lang]["ext"]}')
        else:
            slugs[problem] = os.path.join(d, f'{problem}.{lang_specifics[lang]["ext"]}')

    existing_solutions = list()
    for k, v in slugs.items():
        if os.path.exists(v):
            print(f'Skipping {k}, {v} already exists. ')
            existing_solutions.append(k)
    for i in existing_solutions:
        del slugs[i]

    return slugs


def get_slug_data(name: str) -> dict:
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


def get_test_cases(data: dict) -> list:
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


def generate_test_code(tests: list, code: str) -> str:
    tester = '\n\nfunc main() {\n'

    if '/**' in code:
        code = code.split('*/\n')[-1]

    func_decl = code.split('\n')[0]
    func_name = func_decl.replace('func ', '').split('(')[0]
    func_params = {p.split(' ')[0]: ''.join(p.split(' ')[1:]) for p in func_decl.split('(')[1].split(')')[0].split(', ')}
    # param_types = [v for v in func_params.values()]
    # same_type = all(p == param_types[0] for p in param_types[1:])
    for i, t in enumerate(tests):
        param_def = f'\t// Example {i + 1} \n'
        args = []
        for vn, vt in func_params.items():
            param_def += f'\tvar {vn}{i + 1} {vt} = {t["input"][vn]}\n'
            args.append(f'{vn}{i + 1}')
        run_def = f'\tfmt.Println("Expected: {t["output"]}\tOutput: ", {func_name}({", ".join(args)}))\n\n'

        tester += param_def + run_def

    return tester + '}\n'


def create_code_template(slug: str, file: str, lang: str, no_tests: bool, data: dict) -> None:
    source_link = f'https://leetcode.com/problems/{slug}/'
    code_snippet = ''

    for snippet in data['code_snippets']:
        if snippet['lang_slug'] == lang:
            code_snippet = snippet['code']
            break

    code = ''
    if 'prefix' in lang_specifics[lang].keys():
        code = lang_specifics[lang]['prefix']

    code += f'{lang_specifics[lang]["com"]} source: {source_link}\n\n'
    code += code_snippet

    # experimental feature, tested only with golang
    if lang == 'golang' and not no_tests:
        code += generate_test_code(get_test_cases(data), code_snippet)

    with open(file, 'w') as code_file:
        code_file.write(code)


def clear_leetcode_meta_file(data: dict):
    # keeping meta.csv actual by removing problems with no solution
    # obviously not the prettiest way to do it
    solutions = dict()
    for f in os.listdir('../leetcode'):
        if not f.startswith('.'):
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
            if lang_specifics[old_lang]['ext'] not in solutions[problem]:
                if problem not in lang_to_del:
                    lang_to_del[problem] = [lang_id]
                else:
                    lang_to_del[problem].append(lang_id)

    for problem, langs in lang_to_del.items():
        for lang_id in langs:
            del data[problem]['lang'][lang_id]

    for problem in problem_to_del:
        del data[problem]

    pd.DataFrame.from_dict(data).transpose().to_csv('../leetcode/.meta.csv', index_label='slug', index=True)


def generate_leetcode_readme(data: dict):
    header = '## leetcode solutions'
    problems_table = '|Problem|Solution|Difficulty|Tags|\n|-|-|-|-|\n'
    stats = {
        'solved': 0,
        'by_diff': dict(),
        'by_tags': dict(),
        'by_lang': dict()
    }
    # Getting simple stats and generating rows for the table
    for problem, meta in data.items():
        stats['solved'] += 1
        if meta['difficulty'] not in stats['by_diff']:
            stats['by_diff'][meta['difficulty']] = 0
        stats['by_diff'][meta['difficulty']] += 1

        for tag in meta['tags']:
            if tag not in stats['by_tags']:
                stats['by_tags'][tag] = 0
            stats['by_tags'][tag] += 1

        for lang in meta['lang']:
            if lang not in stats['by_lang']:
                stats['by_lang'][lang] = 0
            stats['by_lang'][lang] += 1

        solutions_links = ' '.join([f'[{lang}](/leetcode/{problem}.{lang_specifics[lang]["ext"]})' for lang in meta['lang']])
        problems_table += f'| [{meta["id"]}. {problem.replace("-", " ").capitalize() }](https://leetcode.com/problems/{problem}/) ' \
                          f'| {solutions_links} ' \
                          f'| {meta["difficulty"]} ' \
                          f'| {", ".join(meta["tags"])} |\n'
    line_break = '\n\t'
    tab = ',\t'
    stats_str = f'''```
leetcode statistics:
 {stats['solved']} problems solved in total
 - Solved problems by difficulty:
\t{line_break.join([diff + ' - ' + str(count) for diff, count in stats['by_diff'].items()])}
 - Solved problems by tags:
\t{tab.join([tag + ' - ' + str(count) for tag, count in stats['by_tags'].items()])}
 - Solutions by languages:
\t{line_break.join([lang + ' - ' + str(count) for lang, count in stats['by_lang'].items()])}
```'''
    # readme = header + '\n\n' + stats_str + '\n\n' + problems_table
    readme = header + '\n\n' + problems_table
    with open('../leetcode/README.md', 'w') as readme_file:
        readme_file.write(readme)


def update_meta_file(slug: str, lang: str, slug_data: dict) -> None:
    if not os.path.exists('../leetcode/.meta.csv'):
        with open('../leetcode/.meta.csv', 'w') as f:
            f.write('slug,id,difficulty,tags,lang\n')
    data = pd.read_csv('../leetcode/.meta.csv', index_col='slug', converters={'lang': pd.eval}).to_dict('index')

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

    pd.DataFrame.from_dict(data).transpose().to_csv('../leetcode/.meta.csv', index_label='slug', index=True)


def main():
    parser = argparse.ArgumentParser(description="Pulls problems from leetcode and sets up env in local repo")
    parser.add_argument('problems', default='', type=str, nargs='+', help='Name or link to the problems')
    parser.add_argument('--lang', dest='lang', action='store', default='golang', help='Solution lang')
    parser.add_argument('--dir', dest='dir', action='store', default='../leetcode', help='Solutions directory')
    parser.add_argument('--no-tests', dest='no_tests', action='store_true', default=False, help='Suppress automatic test creation')
    args = parser.parse_args()

    if args.problems != '':
        slugs = get_title_slugs(args.problems, args.lang, args.dir)
        if len(slugs) == 0:
            print("Nothing to do")
            return

        for slug, file in slugs.items():
            slug_data = get_slug_data(slug)['data']['question']
            create_code_template(slug, file, args.lang, args.no_tests, slug_data)
            update_meta_file(slug, args.lang, slug_data)

    data = pd.read_csv('../leetcode/.meta.csv', index_col='slug', converters={'lang': pd.eval, 'tags': pd.eval}).to_dict('index')
    clear_leetcode_meta_file(data)
    generate_leetcode_readme(data)


if __name__ == '__main__':
    main()
