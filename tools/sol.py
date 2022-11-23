import os.path

import argparse
from functools import reduce

DEFAULT_DIR = f"{os.path.dirname(__file__)}/.."
DEFAULT_HANDLE = "dimaglushkov"
allowed_calls = {
    "codeforces": {"pull", "delete"},
    "leetcode":   {"pull", "delete"},
    "contests":   {"pre", "post"},
    "readme":     {"generate"},
}

target_aliases = {
    "cf":  "codeforces",
    "lc":  "leetcode",
    "c":   "contests",
    "con": "contests",
    "rm":  "readme"
}
action_aliases = {
    "p": "pull",
    "d": "delete",
    "g": "generate",
}

def main():
    all_targets = list(set(list(allowed_calls.keys()) + list(reduce(lambda x, y: x + y, target_aliases.items()))))
    all_actions = list()
    for i in allowed_calls.values():
        for j in i:
            all_actions.append(j)
    all_actions = list(set(all_actions + list(reduce(lambda x, y: x + y, action_aliases.items()))))

    parser = argparse.ArgumentParser("Simple tool to manage solutions")
    parser.add_argument("target", action="store", type=str, choices=all_targets)
    parser.add_argument("action", action="store", type=str, choices=all_actions)
    parser.add_argument('--lang', dest='lang', action='store', default='golang', help='Solution language')
    parser.add_argument('--dir', dest='dir', action='store', default=DEFAULT_DIR, help='Solutions directory')
    parser.add_argument('--handle', dest='handle', action='store', default=DEFAULT_HANDLE, help='Codeforces handle')
    parser.add_argument('val', default='', type=str, nargs='+', help='Values for a chosen call')

    args = parser.parse_args()

    target = target_aliases[args.target] if args.target in target_aliases else args.target
    action = action_aliases[args.action] if args.action in action_aliases else args.action

    if action not in allowed_calls[target]:
        print(f"Incorrect target & action combination: {target} {action}")
        return

    if target == "codeforces":
        from lib import codeforces
        t = codeforces
        sol_dir = args.dir if args.dir != DEFAULT_DIR else os.path.join(DEFAULT_DIR, target)

    elif target == "leetcode":
        from lib import leetcode
        t = leetcode
        sol_dir = args.dir if args.dir != DEFAULT_DIR else os.path.join(DEFAULT_DIR, target)

    elif target == "contests":
        from lib import contests
        t = contests
        sol_dir = args.dir if args.dir != DEFAULT_DIR else os.path.join(DEFAULT_DIR, target)

    elif target == "readme":
        from lib import readme
        t = readme
        sol_dir = args.dir

    else:
        print(f"Unknown target: {target}")
        return

    if action == "pull":
        t.pull(args.val, args.lang, sol_dir)
    elif action == "delete":
        t.delete(args.val, args.lang, sol_dir)
    elif action == "generate":
        t.generate(sol_dir)
    elif action == "pre":
        t.pre(args.val, args.lang, sol_dir)
    elif action == "post":
        t.post(args.val, args.handle, sol_dir)
    elif action == "stats":
        t.stats(sol_dir)
    else:
        print("Unknown action")


if __name__ == '__main__':
    main()
