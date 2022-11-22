import os

import pandas as pd
import requests

from . import shared

CHARTS = ['by_difficulty', 'by_tags']
URL_TEMPLATE = 'https://codeforces.com/contest/{}/problem/{}'
META_FILE = '.meta.csv'
TEMPLATES_DIR = shared.get_templates_dir()
LANG_SPECS = shared.get_lang_specs()


def _generate_codeforces_readme(data: dict, sol_dir: str):
    header = '## codeforces'
    problems_table = '''| Problem | Solution | Difficulty | Tags |
|-|-|-|-|
'''
    for k, v in data.items():
        solutions_links = ', '.join(
            [f'[{lang}](/codeforces/{k}/{k}.{LANG_SPECS[lang]["ext"]})' for lang in v['lang']])

        problems_table += \
            f'| [{k}. {v["name"]}](https://codeforces.com/contest/{v["contest"]}/problem/{v["problem"]}) ' \
            f'| {solutions_links} ' \
            f'| {str(v["difficulty"]).replace(".0", "").replace("nan", "")} ' \
            f'| {", ".join([tag for tag in v["tags"] if str(v["difficulty"]) not in tag])} |\n'

    stats_str = shared.generate_svg_stats(data, sol_dir, CHARTS)
    readme = header + stats_str + problems_table
    with open(os.path.join(sol_dir, 'README.md'), 'w') as readme_file:
        readme_file.write(readme)


def _get_problems_id(problems: list) -> list:
    ids = list(list())
    for problem in problems:
        if not problem.startswith('http'):
            print(f"Please, provide link to the problem: {problem}")
            continue
        contest_id = problem.split('/')[-3]
        index = problem.split('/')[-1]
        ids += [[int(contest_id), index]]

    return ids


def _get_problem_data(problem_id: list) -> dict:
    all_problems_data = requests.get("https://codeforces.com/api/problemset.problems?locale=en").json()
    if all_problems_data["status"] != "OK":
        print("Not Ok response")
        exit(1)

    all_problems_data = all_problems_data["result"]["problems"]

    def search(contest_id, _problem_id):
        i = 0
        while not (all_problems_data[i]["contestId"] == contest_id and all_problems_data[i]["index"] == _problem_id):
            i += 1
        return i

    data_problem_id = search(problem_id[0], problem_id[1])
    return all_problems_data[data_problem_id]


def _clear_codeforces_meta_file(data: dict, sol_dir: str):
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

    pd.DataFrame.from_dict(data).transpose().to_csv(os.path.join(sol_dir, META_FILE), index_label='id', index=True)


def _create_code_template(data: dict, lang: str, sol_dir: str) -> None:
    problem_id = f"{data['contestId']}{data['index']}"
    if not os.path.exists(os.path.join(sol_dir, problem_id)):
        os.mkdir(os.path.join(sol_dir, problem_id))
    file = os.path.join(sol_dir, problem_id, f"{problem_id}.{LANG_SPECS[lang]['ext']}")

    source_link = URL_TEMPLATE.format(data["contestId"], data["index"])
    with open(os.path.join(TEMPLATES_DIR, f"codeforces.{LANG_SPECS[lang]['ext']}")) as f:
        template = f.readlines()
    i = 0
    while i < len(template):
        if "source: _" in template[i]:
            break
        else:
            i += 1
    template[i] = template[i].replace("_", source_link)
    code = "".join(template)

    with open(file, 'w') as code_file:
        code_file.write(code)


def _update_meta_file(problem_data: dict, lang: str, sol_dir: str) -> None:
    meta_file = os.path.join(sol_dir, META_FILE)
    if not os.path.exists(meta_file):
        with open(meta_file, 'w') as f:
            f.write('id,name,contest,problem,difficulty,tags,lang\n')
    data = pd.read_csv(meta_file, index_col="id", converters={'lang': pd.eval}).to_dict('index')

    # adding new problem or solution
    problem_id = f"{problem_data['contestId']}{problem_data['index']}"
    if problem_id not in data.keys():
        data[problem_id] = {
            'name': problem_data['name'],
            'contest': problem_data['contestId'],
            'problem': problem_data['index'],
            'difficulty': str(problem_data['rating']) if "rating" in problem_data else "",
            'tags': problem_data['tags'] if "tags" in problem_data else [],
            'lang': [lang]
        }
    else:
        data[problem_id]['lang'].append(lang)

    df = pd.DataFrame.from_dict(data).transpose()
    df.to_csv(meta_file, index_label='id', index=True, float_format='%.0f')


def pull(problems: list, lang: str, sol_dir: str):
    if lang not in LANG_SPECS:
        print(f"Language {lang} is not supported yet\n"
              "In order to add it you need to add its extension to tools/lib/cf-problem-puller.py, "
              "look for LANG_SPECIFICS "
              "& create corresponding template file in tools/templates.")
        return

    if problems != ['.'] and problems != []:
        ids = _get_problems_id(problems)

        # skip existing solutions
        temp = list()
        for i in ids:
            problem_id = f"{i[0]}{i[1]}"
            if not os.path.exists(os.path.join(sol_dir, problem_id, f"{problem_id}.{LANG_SPECS[lang]['ext']}")):
                temp += [i]
        ids = temp

        if len(ids) == 0:
            print("Nothing to do")
            return

        for problem_id in ids:
            data = _get_problem_data(problem_id)
            _create_code_template(data, lang, sol_dir)
            _update_meta_file(data, lang, sol_dir)

    data = pd.read_csv(os.path.join(sol_dir, META_FILE), index_col='id',
                       converters={'lang': pd.eval, 'tags': pd.eval}).to_dict('index')
    _clear_codeforces_meta_file(data, sol_dir)
    _generate_codeforces_readme(data, sol_dir)


def delete(problems: list, lang: str, sol_dir: str):
    ids = _get_problems_id(problems)
    cnt = len(ids)
    for i in ids:
        problem_id = f"{i[0]}{i[1]}"
        solution_path = os.path.join(sol_dir, problem_id, f"{problem_id}.{LANG_SPECS[lang]['ext']}")
        if not os.path.exists(solution_path):
            cnt -= 1
            continue

        os.remove(solution_path)
        os.rmdir(os.path.join(sol_dir, problem_id))

    if cnt == 0:
        print("Nothing to do")
        return

    data = pd.read_csv(os.path.join(sol_dir, META_FILE), index_col='id',
                       converters={'lang': pd.eval, 'tags': pd.eval}).to_dict('index')
    _clear_codeforces_meta_file(data, sol_dir)
    _generate_codeforces_readme(data, sol_dir)
