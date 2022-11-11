import os.path

import argparse

DEFAULT_DIR = f"{os.path.dirname(__file__)}/.."

def main():
    aliases = {
        "cf": "codeforces",
        "lc": "leetcode",
        "rm": "readme",
        "p": "pull",
        "d": "delete",
        "g": "generate"
    }

    parser = argparse.ArgumentParser("Simple tool to manage solutions")
    parser.add_argument(
        "target", action="store", type=str, choices=["cf", "codeforces", "lc", "leetcode", "rm", "readme"]
    )
    parser.add_argument(
        "action", action="store", type=str, choices=["p", "pull", "d", "delete", "g", "generate"]
    )
    parser.add_argument('--lang', dest='lang', action='store', default='golang', help='Solution language')
    parser.add_argument('--dir', dest='dir', action='store', default=DEFAULT_DIR, help='Solutions directory')
    parser.add_argument('problems', default='', type=str, nargs='+', help='Names or links of problems')

    args = parser.parse_args()

    target = aliases[args.target] if args.target in aliases else args.target
    action = aliases[args.action] if args.action in aliases else args.action

    if target == "codeforces" and (action == "pull" or action == "delete"):
        from lib import codeforces
        t = codeforces
        sol_dir = args.dir if args.dir != DEFAULT_DIR else os.path.join(DEFAULT_DIR, target)

    elif target == "leetcode" and (action == "pull" or action == "delete"):
        from lib import leetcode
        t = leetcode
        sol_dir = args.dir if args.dir != DEFAULT_DIR else os.path.join(DEFAULT_DIR, target)

    elif target == "readme" and action == "generate":
        from lib import readme
        t = readme
        sol_dir = args.dir

    else:
        print(f"Incorrect target & action combination: {target} {action}")
        return

    if action == "pull":
        t.pull(args.problems, args.lang, sol_dir)
    elif action == "delete":
        t.delete(args.problems, args.lang, sol_dir)
    elif action == "generate":
        t.generate(sol_dir)
    else:
        print("Unknown action")


if __name__ == '__main__':
    main()
