import datetime
import json
import os
import shutil

import requests

from . import shared, codeforces, leetcode

TEMPLATES_DIR = shared.get_templates_dir()
LANG_SPECS = shared.get_lang_specs()
CF_PROBLEM_LIST = ['A', 'B', 'C', 'D', 'E', 'F', 'G']
LC_PROBLEM_LIST = [1, 2, 3, 4]
DATE_FORMAT = '%-d %b %Y'
DEFAULT_USERNAME = 'dimaglushkov'


def run(values, lang, sol_dir, handle, action):
    if action == 'pre':
        pre(values, lang, sol_dir)
    elif action == 'post':
        post(values, handle, sol_dir)
    elif action == 'pull':
        pull(values, lang, sol_dir)
    elif action == 'stats':
        stats(handle, sol_dir)


def _cf_get_contest_meta(url: str) -> dict:
    res = dict()
    contests = requests.get('https://codeforces.com/api/contest.list').json()['result']
    if url.endswith('/'):
        url = url.rstrip('/')

    idx = int(url.split('/')[-1])
    name = ''
    for i in contests:
        if i['id'] == idx:
            name = i['name']
            break
    if name == '':
        raise ValueError(f'Couldn\'t find contest with id {idx}')

    ext_id = ''
    if '#' in name:
        ext_id = name.split('#')[1].split(' ')[0]
    else:
        ext_id = name.split('Round ')[1].split(' ')[0]

    res['name'] = name
    res['ext_id'] = ext_id
    res['short_name'] = f'cf-round-{ext_id}'
    return res


def _pull_codeforces(url: str, lang: str, sol_dir: str):
    raise ValueError("Not implemented")


def _pre_codeforces(url: str, lang: str, sol_dir: str):
    meta = _cf_get_contest_meta(url)
    name = meta['name']
    short_name = meta['short_name']

    contest_sol_dir = os.path.join(sol_dir, short_name)
    if os.path.isdir(contest_sol_dir):
        raise ValueError(f'Directory {contest_sol_dir} already exists')

    # Generating solution templates
    os.mkdir(contest_sol_dir)
    with open(os.path.join(TEMPLATES_DIR, f"codeforces.{LANG_SPECS[lang]['ext']}"), 'r') as templ_file:
        templ_data = templ_file.read()

        for i in CF_PROBLEM_LIST:
            os.mkdir(os.path.join(contest_sol_dir, i))
            code_file_path = os.path.join(contest_sol_dir, i, f"{i}.{LANG_SPECS[lang]['ext']}")
            with open(code_file_path, "w") as file:
                file.write(
                    templ_data.replace(
                        'source: _', f"source: {url.replace('contests', 'contest')}/problem/{i}"
                    )
                )
            print(f'{i}: file://{code_file_path}')

    # Updating README
    with open(os.path.join(sol_dir, 'README.md'), 'a') as readme_file:
        date_str = datetime.datetime.now().strftime(DATE_FORMAT).lower()
        contests_url = url
        if 'contests' not in contests_url:
            contests_url = contests_url.replace('contest', 'contests')
        readme_file.write(
            f'| [{name}]({contests_url}) | ? / ? | [solutions](/contests/{short_name}) | {date_str} |\n'
        )


def _post_codeforces(url: str, handle: str, sol_dir: str):
    cf_sol_path = os.path.abspath(os.path.join(sol_dir, '..', 'codeforces'))
    if not os.path.isdir(cf_sol_path):
        print(f'{cf_sol_path} doesnt exist')
        return

    meta = _cf_get_contest_meta(url)

    if url.endswith('/'):
        url = url.rstrip('/')
    idx = int(url.split('/')[-1])

    data = requests.get(f'https://codeforces.com/api/contest.status?contestId={idx}&handle={handle}').json()
    if data['status'] != 'OK':
        print('Not OK response')
        return

    solved_problems = dict()
    for i in data['result']:
        if i['verdict'] == 'OK':
            pid = i['problem']['index']

            solved_problems[pid] = i['problem']
            lang = i['programmingLanguage']

            if lang.lower() == 'go':
                lang = 'golang'
            solved_problems[pid]['lang'] = lang

    sp_list = sorted([v for v in solved_problems.values()], key=lambda x: x['index'])
    for p in sp_list:
        codeforces.update_meta_file(p, p['lang'], cf_sol_path)
        full_id = f"{p['contestId']}{p['index']}"
        os.mkdir(os.path.join(cf_sol_path, full_id))
        shutil.copyfile(
            os.path.join(sol_dir, meta['short_name'], p['index'], f"{p['index']}.{LANG_SPECS[p['lang']]['ext']}"),
            os.path.join(cf_sol_path, full_id, f"{full_id}.{LANG_SPECS[p['lang']]['ext']}")
        )
    codeforces.update_codeforces_readme(cf_sol_path)


def _lc_get_contest_meta(contest_url: str) -> (bool, list):
    response = requests.get(contest_url.replace("/contest/", "/contest/api/info/"))
    if response.status_code != 200:
        raise ValueError("Not OK response for url: " + contest_url)
    data = response.json()

    if data["contest"]["start_time"] > data["current_timestamp"]:
        return False, None

    return True, [q["title_slug"] for q in data["questions"]]


def _pull_leetcode(url: str, lang: str, sol_dir: str):
    if url.endswith("/"):
        url = url.rstrip("/")
    tag = url.split("/")[-1]
    con_type = tag.split('-')[0]
    con_idx = tag.split('-')[-1]
    short_name = f"lc-{con_type}-{con_idx}"
    contest_sol_dir = os.path.join(sol_dir, short_name)
    if os.path.isdir(contest_sol_dir):
        print(f"Solutions directory {contest_sol_dir} already exists")
        return
    started, question_slugs = _lc_get_contest_meta(url)
    if not started:
        print("Contest has not started yet")
        return

    os.mkdir(contest_sol_dir)
    for i, slug in enumerate(question_slugs):
        problem_sol_dir = os.path.join(contest_sol_dir, str(i + 1))
        file = os.path.join(problem_sol_dir, f"{i + 1}.{LANG_SPECS[lang]['ext']}")
        os.mkdir(problem_sol_dir)

        slug_data = leetcode.get_slug_data(slug)['data']['question']
        if slug_data is not None:
            leetcode.create_code_template(slug, file, lang, slug_data)
        else:
            print("Cannot pull problem meta data")
            exit(1)
        print(f"{i + 1}: file://{os.path.abspath(file)}")

    with open(os.path.join(sol_dir, "README.md"), "a") as readme_file:
        name = f"Leetcode {con_type.capitalize()} Contest {con_idx}"
        date_str = datetime.datetime.now().strftime(DATE_FORMAT).lower()
        readme_file.write(
            f"| [{name}]({url}) | ? / ? | [solutions](/contests/{short_name}) | {date_str} |\n"
        )


def _pre_leetcode(url: str, lang: str, sol_dir: str):
    if url.endswith("/"):
        url = url.rstrip("/")
    tag = url.split("/")[-1]
    con_type = tag.split('-')[0]
    con_idx = tag.split('-')[-1]
    short_name = f"lc-{con_type}-{con_idx}"
    contest_sol_dir = os.path.join(sol_dir, short_name)
    if os.path.isdir(contest_sol_dir):
        print(f"Solutions directory {contest_sol_dir} already exists")
        return
    _, question_slugs = _lc_get_contest_meta(url)

    os.mkdir(contest_sol_dir)
    for i in LC_PROBLEM_LIST:
        problem_sol_dir = os.path.join(contest_sol_dir, str(i + 1))
        os.mkdir(problem_sol_dir)

        file = os.path.join(problem_sol_dir, f"{i}.{LANG_SPECS[lang]['ext']}")
        shutil.copyfile(os.path.join(TEMPLATES_DIR, f"leetcode.{LANG_SPECS[lang]['ext']}"), file)
        print(f'{i}: file://{os.path.abspath(file)}')

    with open(os.path.join(sol_dir, "README.md"), "a") as readme_file:
        name = f"Leetcode {con_type.capitalize()} Contest {con_idx}"
        date_str = datetime.datetime.now().strftime(DATE_FORMAT).lower()
        readme_file.write(
            f"| [{name}]({url}) | ? / ? | [solutions](/contests/{short_name}) | {date_str} |\n"
        )


def _post_leetcode(url, sol_dir):
    ok, questions = _lc_get_contest_meta(contest_url=url)
    if not ok:
        print("Not OK response for contest meta")
        return

    solved = [False, False, False, False]

    for i, q in enumerate(questions):
        x = leetcode.get_slug_data(q)


def _gen_stats_json(sol_dir: str, platform: str, data: dict):
    data["updated_at"] = datetime.datetime.now().strftime(DATE_FORMAT)
    with open(os.path.join(sol_dir, f".{platform}_stats.json"), "w") as f:
        f.write(json.dumps(data))


def _gen_lc_rating_stats(sol_dir: str, username: str):
    ses = requests.session()
    ses.get("https://leetcode.com/")
    ses.headers["Referer"] = "https://leetcode.com/dimaglushkov/"
    ses.headers["Content-Type"] = "application/json"
    if 'csrftoken' in ses.cookies:
        ses.headers["x-csrftoken"] = ses.cookies['csrftoken']
    else:
        ses.headers["x-csrftoken"] = ses.cookies['csrf']

    query = """
query userContestRankingInfo($username: String!) {
  userContestRanking(username: $username) {
    attendedContestsCount
    rating
    globalRanking
    totalParticipants
    topPercentage
    badge {
      name
    }
  }
  userContestRankingHistory(username: $username) {
    attended
    trendDirection
    problemsSolved
    totalProblems
    finishTimeInSeconds
    rating
    ranking
    contest {
      title
      startTime
    }
  }
}
    """

    resp = ses.post(
        url="https://leetcode.com/graphql/",
        json={
            "query": query,
            "variables": {
                "username": username
            }
        }
    )
    if resp.status_code != 200:
        print("Not 200 response status code")
        print(resp)
        exit(1)

    data = resp.json()["data"]
    contests = [c for c in data["userContestRankingHistory"] if c["attended"]]
    dates, ratings = list(), list()
    for c in contests:
        dates.append(datetime.datetime.fromtimestamp(c["contest"]["startTime"]))
        ratings.append(int(c["rating"]))

    # Generate SVG chart with rating changes
    shared.generate_date_based_plot(os.path.join(sol_dir, ".leetcode_rating.svg"), dates, ratings)

    # Generate stats JSON
    _gen_stats_json(
        sol_dir=sol_dir,
        platform="leetcode",
        data={
            "attended": data["userContestRanking"]["attendedContestsCount"],
            "rating": int(data["userContestRanking"]["rating"]),
            "top": data["userContestRanking"]["topPercentage"]
        })


def _gen_cf_rating_stats(sol_dir: str, username: str):
    ses = requests.session()
    resp = ses.get(f"https://codeforces.com/api/user.rating?handle={username}")
    if resp.status_code != 200:
        print("Not 200 response status code")
        print(resp)
        exit(1)

    contests = resp.json()
    if contests["status"] != "OK":
        print("Not OK response status")
        exit(1)
    contests = contests["result"]
    dates, ratings = list(), list()
    for c in contests:
        # for cf using ratingUpdateTimeSeconds instead of actual contest time just because it's easier to implement
        dates.append(datetime.datetime.fromtimestamp(c["ratingUpdateTimeSeconds"]))
        ratings.append(int(c["newRating"]))

    # Generate SVG chart with rating changes
    shared.generate_date_based_plot(os.path.join(sol_dir, ".codeforces_rating.svg"), dates, ratings)

    # Generate stats JSON
    rating_list_resp = ses.get("https://codeforces.com/api/user.ratedList?activeOnly=true&includeRetired=false")
    if rating_list_resp.status_code != 200:
        print("Not 200 response status code")
        print(rating_list_resp)
        exit(1)
    rating_list = rating_list_resp.json()
    if rating_list["status"] != "OK":
        print("Not OK response status")
        exit(1)
    cur_rating = ratings[-1]
    pos, total = 1, 1
    rating_list = [r["rating"] for r in rating_list["result"]]
    total = len(rating_list)
    for i, r in enumerate(rating_list):
        if r == cur_rating:
            pos = i
            break
    _gen_stats_json(
        sol_dir=sol_dir,
        platform="codeforces",
        data={
            "attended": len(ratings),
            "rating": cur_rating,
            "top": round(pos / total * 100, 2)
        })


def stats(handle: str, sol_dir: str):
    _gen_lc_rating_stats(sol_dir, handle)
    _gen_cf_rating_stats(sol_dir, handle)

    dates = list()
    top = list()
    with open(os.path.join(sol_dir, "README.md"), "r") as readme_file:
        lines = readme_file.readlines()
        for i in lines:
            fields = i.split('|')
            if len(fields) < 5 or '?' in fields[2] or '/' not in fields[2]:
                continue

            res = fields[2].split(' / ')
            me, total = int(res[0]), int(res[1])

            top.append(round(me / total * 100, 2))
            dates.append(datetime.datetime.strptime(fields[4].title().strip(), DATE_FORMAT.replace("-", "")))

    shared.generate_date_based_plot(os.path.join(sol_dir, ".stats.svg"), dates, top)


# pull tries to pull problem-related data and build templates based on it
# however, it usually doesn't work for the ongoing contests
def pull(vals: list, lang: str, sol_dir: str):
    if len(vals) < 1:
        return
    url = vals[0]
    if "codeforces" in url:
        _pull_codeforces(url, lang, sol_dir)
    elif "leetcode" in url:
        _pull_leetcode(url, lang, sol_dir)


# pre generates solutions from based on the template content and update contests list
def pre(vals: list, lang: str, sol_dir: str):
    if len(vals) < 1:
        return
    url = vals[0]
    if "codeforces" in url:
        _pre_codeforces(url, lang, sol_dir)
    elif "leetcode" in url:
        _pre_leetcode(url, lang, sol_dir)


def post(vals: list, handle: str, sol_dir: str):
    if len(vals) < 1:
        return
    url = vals[0]
    if "codeforces" in url:
        _post_codeforces(url, handle, sol_dir)
    elif "leetcode" in url:
        _post_leetcode(url, sol_dir)
