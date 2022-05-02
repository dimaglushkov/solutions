import os
import shutil

import pandas as pd


def move_files():
    solutions = [s for s in os.listdir("../codeforces") if s.endswith(".py")]
    dirs = [s.split(".py")[0] for s in solutions]
    for d, s in zip(dirs, solutions):
        os.mkdir("../codeforces/" + d)
        shutil.move(os.path.join("../codeforces/", s), os.path.join("../codeforces", d, s))


def main():
    data = pd.read_csv('../codeforces/.meta.csv', index_col='slug', converters={ 'tags': pd.eval}).to_dict('index')
    transformed_data = dict()
    # f.write('id,name,contest,problem,difficulty,tags,lang\n')

    for k, v in data.items():
        ind = f"{v['contest']}{v['problem']}"
        transformed_data[ind] = {
            'name': k,
            'contest': v['contest'],
            'problem': v['problem'],
            'difficulty': v['difficulty'],
            'tags': v['tags'],
            'lang': [v['lang']],
        }

    pd.DataFrame.from_dict(transformed_data).transpose().to_csv('../codeforces/.new_meta.csv', index_label='id', index=True)


if __name__ == '__main__':
    main()
