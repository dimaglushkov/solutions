import argparse
import os

import pandas as pd
import requests

url_template = 'https://codeforces.com/contest/{}/problem/{}'

lang_specifics = {
    'golang': {
        'ext': 'go',
        'com': '//'
    },
    'python': {
        'ext': 'py',
        'com': '#',
    }
}


def generate_codeforces_readme(data: dict):
    text = '''## codeforces
| Problem | Solution | Difficulty | Tags |
|-|-|-|-|
'''

    with open('../../codeforces/README.md', 'w') as f:
        for k, v in data.items():
            solutions_links = ', '.join(
                [f'[{lang}](/codeforces/{k}/{k}.{lang_specifics[lang]["ext"]})' for lang in v['lang']])

            text += f'| [{k}. {v["name"]}](https://codeforces.com/contest/{v["contest"]}/problem/{v["problem"]}) ' \
                    f'| {solutions_links} ' \
                    f'| {str(v["difficulty"]).replace(".0", "").replace("nan", "")} ' \
                    f'| {", ".join([tag for tag in v["tags"] if str(v["difficulty"]) not in tag])} |\n'
        f.write(text)


def get_problems_id(problems: list, lang: str, d: str) -> list:
    ids = list(list())
    for problem in problems:
        if not problem.startswith('http'):
            print(f"Please, provide a link to the problem: {problem}")
        contest_id = problem.split('/')[-3]
        index = problem.split('/')[-1]

        problem_id = f"{contest_id}{index}"
        if not os.path.exists(os.path.join(d, problem_id, f"{problem_id}.{lang_specifics[lang]['ext']}")):
            ids += [[int(contest_id), index]]

    return ids


def get_problem_data(problem_id: list) -> dict:
    all_problems_data = requests.get("https://codeforces.com/api/problemset.problems?locale=en").json()
    if all_problems_data["status"] != "OK":
        print("Not Ok response")
        exit(1)

    all_problems_data = all_problems_data["result"]["problems"]

    def search(contest_id, problem_id):
        i = 0
        while not (all_problems_data[i]["contestId"] == contest_id and all_problems_data[i]["index"] == problem_id):
            i += 1
        return i

    data_problem_id = search(problem_id[0], problem_id[1])
    return all_problems_data[data_problem_id]


def clear_codeforces_meta_file(data: dict):
    # keeping meta.csv actual by removing problems with no solution
    # obviously not the prettiest way to do it
    solutions = dict()
    for pkg in os.listdir('../../codeforces'):
        if os.path.isdir(f'../codeforces/{pkg}'):
            for f in os.listdir(f'../codeforces/{pkg}'):
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

    pd.DataFrame.from_dict(data).transpose().to_csv('../../codeforces/.meta.csv', index_label='id', index=True)


def create_code_template(lang: str, d: str, no_tests: bool, data: dict) -> None:
    problem_id = f"{data['contestId']}{data['index']}"
    if not os.path.exists(os.path.join(d, problem_id)):
        os.mkdir(os.path.join(d, problem_id))
    file = os.path.join(d, problem_id, f"{problem_id}.{lang_specifics[lang]['ext']}")

    source_link = url_template.format(data["contestId"], data["index"])
    code = ""

    if lang == "golang":
        template = ""
        with open(os.path.join(d, ".template.go")) as f:
            template = f.readlines()
        i = 0
        while i < len(template):
            if "func solve(" in template[i]:
                break
            else:
                i += 1

        template.insert(i, f"\n// source: {source_link}\n")
        code = "".join(template)
    else:
        if "prefix" in lang_specifics[lang]:
            code += lang_specifics[lang]["prefix"]
        code += f"{lang_specifics[lang]['com']} source: {source_link}\n\n"
        code += lang_specifics[lang]["main"]

    # if not no_tests and (lang == "golang" or lang == "go"):
    #     generate_tests()

    with open(file, 'w') as code_file:
        code_file.write(code)


def update_meta_file(lang: str, d: str, problem_data: dict) -> None:
    meta_file = os.path.join(d, ".meta.csv")
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

def main():
    parser = argparse.ArgumentParser(description="Pulls problems from codeforces and sets up env in local repo")
    parser.add_argument('problems', default='', type=str, nargs='+', help='Name or link to the problems')
    parser.add_argument('--lang', dest='lang', action='store', default='golang', help='Solution lang')
    parser.add_argument('--dir', dest='dir', action='store', default='../codeforces', help='Solutions directory')
    parser.add_argument('--no-tests', dest='no_tests', action='store_true', default=False, help='Suppress automatic test creation')
    args = parser.parse_args()

    if args.problems != ['.']:
        ids = get_problems_id(args.problems, args.lang, args.dir)

        if len(ids) == 0:
            print("Nothing to do")
            return

        for problem_id in ids:
            data = get_problem_data(problem_id)
            create_code_template(args.lang, args.dir, args.no_tests, data)
            update_meta_file(args.lang, args.dir, data)

    data = pd.read_csv('../../codeforces/.meta.csv', index_col='id', converters={'lang': pd.eval, 'tags': pd.eval}).to_dict('index')
    clear_codeforces_meta_file(data)
    generate_codeforces_readme(data)


if __name__ == '__main__':
    main()
