import os

import pandas as pd

order = ['contests', 'leetcode', 'codeforces', 'acm.timus']


def main():
    text = ''

    with open('.README.gen.md') as f:
        text += f.read() + '\n\n'

    text += '\n ## Table of Contents\n'
    i = 1
    for d in order:
        text += f'[{i}. {d}](#{d})\n'
        i += 1

    text += '\n\n'

    for d in order:
        readme = f'../{d}/README.md'
        if not os.path.exists(readme):
            print(f'Warning: {readme} does not exist')
            continue
        with open(readme) as f:
            text += f.read() + '\n\n'

    with open('../README.md', 'w') as file:
        file.write(text)


if __name__ == '__main__':
    main()
