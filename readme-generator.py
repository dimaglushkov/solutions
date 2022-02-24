import os

import pandas as pd

order = ['leetcode', 'codeforces', 'acm.timus']


def main():
    text = ''

    with open('.README.gen.md') as f:
        text += f.read() + '\n\n'

    for d in order:
        readme = f'{d}/README.md'
        if not os.path.exists(readme):
            print(f'Warning: {readme} does not exist')
            continue
        with open(readme) as f:
            text += f.read() + '\n\n'

    with open('README.md', 'w') as file:
        file.write(text)


if __name__ == '__main__':
    main()
