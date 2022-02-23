import os

import pandas as pd

lang_specifics = {
    'golang': {
        'ext': 'go',
        'com': '//',
        'prefix': 'package main\n\n'
    },
    'python': {
        'ext': 'py',
        'com': '#'
    }
}


def clear_leetcode_meta_file(data: dict):
    # keeping meta.csv actual by removing problems with no solution
    # obviously not the prettiest way to do it
    solutions = dict()
    for f in os.listdir('leetcode'):
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

    pd.DataFrame.from_dict(data).transpose().to_csv('leetcode/.meta.csv', index_label='slug', index=True)


def generate_leetcode_block(data: dict) -> str:
    table_data = dict()
    header = '## leetcode solutions'
    problems_table = '|Problem|Solutions|Difficulty|Tags|\n|-|-|-|-|\n'
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
        problems_table += f'| [{meta["id"]}. {problem}](https://leetcode.com/problems/{problem}/) ' \
                          f'| {solutions_links} ' \
                          f'| {meta["difficulty"]} ' \
                          f'| {", ".join(meta["tags"])} |\n'
    line_break = '\n\t'
    tab = ',\t'
    stats_str = f'''```
Some simple statistics:
 {stats['solved']} problems solved in total
 - Solved problems by difficulty:
\t{line_break.join([diff + ' - ' + str(count) for diff, count in stats['by_diff'].items()])}
 - Solved problems by tags:
\t{tab.join([tag + ' - ' + str(count) for tag, count in stats['by_tags'].items()])}
 - Solutions by languages:
\t{line_break.join([lang + ' - ' + str(count) for lang, count in stats['by_lang'].items()])}
```'''

    return header + '\n\n' + stats_str + '\n\n' + problems_table


def main():
    text = ''
    leetcode_metadata = pd.read_csv('leetcode/.meta.csv',
                                    index_col='slug',
                                    converters={'tags': pd.eval, 'lang': pd.eval}
                                    ).to_dict('index')
    clear_leetcode_meta_file(leetcode_metadata)
    text += generate_leetcode_block(leetcode_metadata)

    with open('README.md', 'w') as file:
        file.write(text)


if __name__ == '__main__':
    main()
