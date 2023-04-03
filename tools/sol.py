import argparse
import os.path
from functools import reduce

DEFAULT_DIR = os.path.abspath(f'{os.path.dirname(__file__)}/..')
DEFAULT_HANDLE = 'dimaglushkov'

allowed_calls = {
    'codeforces': {
        'pull',
        'delete'
    },
    'leetcode': {
        'pull',
        'delete'
    },
    'contests': {
        'pull',
        'pre',
        'post',
        'stats'
    },
    'readme': {
        'generate'
    },
}

target_aliases = {
    'cf': 'codeforces',
    'lc': 'leetcode',
    'c': 'contests',
    'con': 'contests',
    'md': 'readme',
}
action_aliases = {
    'p': 'pull',
    'd': 'delete',
    'g': 'generate',
}


def main():
    # extract all possible targets and actions and their aliases
    all_targets = list(set(list(allowed_calls.keys()) + list(reduce(lambda x, y: x + y, target_aliases.items()))))
    all_actions = list(
        set([j for i in allowed_calls.values() for j in i]
            + list(reduce(lambda x, y: x + y, action_aliases.items())))
    )

    # parse cli arguments
    parser = argparse.ArgumentParser('Simple tool to manage solutions')
    parser.add_argument('target', action='store', type=str, choices=all_targets)
    parser.add_argument('action', action='store', type=str, choices=all_actions)
    parser.add_argument('--lang', dest='lang', action='store', default='golang', help='Solution language')
    parser.add_argument('--dir', dest='dir', action='store', default=DEFAULT_DIR, help='Solutions directory')
    parser.add_argument('--handle', dest='handle', action='store', default=DEFAULT_HANDLE, help='User\'s handle')
    parser.add_argument('val', default='', type=str, nargs='*', help='Values for a chosen call')
    args = parser.parse_args()
    target = target_aliases[args.target] if args.target in target_aliases else args.target
    action = action_aliases[args.action] if args.action in action_aliases else args.action

    # validate the given target & action pair
    if target not in all_targets or action not in allowed_calls[target]:
        print(f'Incorrect target & action combination: {target} {action}')
        return 1

    # determine actual sol_dir and call the corresponding runner
    sol_dir = args.dir if args.dir != DEFAULT_DIR or target == 'readme' else os.path.join(DEFAULT_DIR, target)
    if target == 'codeforces':
        from lib import codeforces
        t = codeforces
    elif target == 'leetcode':
        from lib import leetcode
        t = leetcode
    elif target == 'contests':
        from lib import contests
        t = contests
    elif target == 'readme':
        from lib import readme
        t = readme
    else:
        return 1

    t.run(args.val, args.lang, sol_dir, args.handle, action)


if __name__ == '__main__':
    main()
