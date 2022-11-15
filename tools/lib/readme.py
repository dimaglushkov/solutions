import os

order = ['tools', 'contests', 'leetcode', 'codeforces', 'acm.timus']
template_path = os.path.join(os.path.dirname(__file__), '..', 'templates', '.README.gen.md')

def generate(sol_dir: str):
    text = ''

    with open(template_path) as f:
        text += f.read() + '\n\n'

    text += '\n ## Table of Contents\n'
    i = 1
    for d in order:
        text += f'[{i}. {d}](#{d.replace(".", "")}) <br>\n'
        i += 1

    text += '\n\n'

    for d in order:
        readme = os.path.join(sol_dir, d, 'README.md')
        if not os.path.exists(readme):
            print(f'Warning: {readme} does not exist')
            continue
        with open(readme) as f:
            text += f.read() + '\n\n'

    with open(os.path.join(sol_dir, 'README.md'), 'w') as file:
        file.write(text)
