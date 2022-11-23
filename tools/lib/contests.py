import datetime
import os
import shutil

import requests

from . import shared, codeforces

TEMPLATES_DIR = shared.get_templates_dir()
LANG_SPECS = shared.get_lang_specs()
CF_NUM_OF_PROBLEMS = 8
LC_NUM_OF_PROBLEMS = 4
DATE_FORMAT = "%-d %b %Y"


def _cf_get_contest_meta(url: str) -> dict:
    res = dict()
    contests = requests.get("https://codeforces.com/api/contest.list").json()["result"]
    if url.endswith("/"):
        url = url.rstrip("/")

    idx = int(url.split("/")[-1])
    name = ""
    for i in contests:
        if i["id"] == idx:
            name = i["name"]
            break
    if name == "":
        raise ValueError(f"Couldn't find contest with id {idx}")

    ext_id = name.split("#")[1].split(" ")[0]

    res["name"] = name
    res["ext_id"] = ext_id
    res["short_name"] = f"cf-round-{ext_id}"
    return res

def _pre_codeforces(url: str, lang: str, sol_dir: str):
    meta = _cf_get_contest_meta(url)
    name = meta["name"]
    short_name = meta["short_name"]

    contest_sol_dir = os.path.join(sol_dir, short_name)
    if os.path.isdir(contest_sol_dir):
        raise ValueError(f"Directory {contest_sol_dir} already exists")

    # Generating solution templates
    os.mkdir(contest_sol_dir)
    let = 'A'
    with open(os.path.join(TEMPLATES_DIR, f"codeforces.{LANG_SPECS[lang]['ext']}"), "r") as templ_file:
        templ_data = templ_file.read()

        for i in range(int(CF_NUM_OF_PROBLEMS)):
            ix = chr(ord(let) + i)
            os.mkdir(os.path.join(contest_sol_dir, ix))
            with open(os.path.join(contest_sol_dir, ix, f"{ix}.{LANG_SPECS[lang]['ext']}"), "w") as sol_file:
                sol_file.write(
                    templ_data.replace(
                        "source: _", f"source: {url.replace('contests', 'contest')}/problem/{ix}"
                    )
                )

    # Updating README
    with open(os.path.join(sol_dir, "README.md"), "a") as readme_file:
        date_str = datetime.datetime.now().strftime(DATE_FORMAT).lower()
        contests_url = url
        if "contests" not in contests_url:
            contests_url = contests_url.replace("contest", "contests")

        readme_file.write(
            f"| [{name}]({contests_url}) | ? / ? | [solutions](/contests/{short_name}) | {date_str} |\n"
        )


def _post_codeforces(url: str, handle: str, sol_dir: str):
    cf_sol_path = os.path.abspath(os.path.join(sol_dir, "..", "codeforces"))
    if not os.path.isdir(cf_sol_path):
        print(f"{cf_sol_path} doesn't exist")
        return

    meta = _cf_get_contest_meta(url)

    if url.endswith("/"):
        url = url.rstrip("/")
    idx = int(url.split("/")[-1])

    data = requests.get(f"https://codeforces.com/api/contest.status?contestId={idx}&handle={handle}").json()
    if data["status"] != "OK":
        print("Not OK response")
        return

    solved_problems = dict()
    for i in data["result"]:
        if i["verdict"] == "OK":
            pid = i["problem"]["index"]

            solved_problems[pid] = i["problem"]
            lang = i["programmingLanguage"]

            if lang.lower() == "go":
                lang = "golang"
            solved_problems[pid]["lang"] = lang

    sp_list = sorted([v for v in solved_problems.values()], key=lambda x: x['index'])
    for p in sp_list:
        codeforces.update_meta_file(p, p["lang"], cf_sol_path)
        full_id = f'{p["contestId"]}{p["index"]}'
        os.mkdir(os.path.join(cf_sol_path, full_id))
        shutil.copyfile(
            os.path.join(sol_dir, meta['short_name'], p["index"], f"{p['index']}.{LANG_SPECS[p['lang']]['ext']}"),
            os.path.join(cf_sol_path, full_id, f"{full_id}.{LANG_SPECS[p['lang']]['ext']}")
        )
    codeforces.update_codeforces_readme(cf_sol_path)


def _pre_leetcode(url: str, lang: str, sol_dir: str):
    if url.endswith("/"):
        url = url.rstrip("/")
    tag = url.split("/")[-1]
    con_type = tag.split('-')[0]
    con_idx = tag.split('-')[-1]
    short_name = f"lc-{con_type}-{con_idx}"

    contest_sol_dir = os.path.join(sol_dir, short_name)
    if os.path.isdir(contest_sol_dir):
        raise ValueError(f"Directory {contest_sol_dir} already exists")
    os.mkdir(contest_sol_dir)

    for i in range(int(LC_NUM_OF_PROBLEMS)):
        ix = str(i+1)
        os.mkdir(os.path.join(contest_sol_dir, ix))
        shutil.copyfile(
            os.path.join(TEMPLATES_DIR, f"leetcode.{LANG_SPECS[lang]['ext']}"),
            os.path.join(contest_sol_dir, ix, f"{ix}.{LANG_SPECS[lang]['ext']}")
        )

    with open(os.path.join(sol_dir, "README.md"), "a") as readme_file:
        name = f"Leetcode {con_type.capitalize()} Contest {con_idx}"
        date_str = datetime.datetime.now().strftime(DATE_FORMAT).lower()
        readme_file.write(
            f"| [{name}]({url}) | ? / ? | [solutions](/contests/{short_name}) | {date_str} |\n"
        )


def _post_leetcode():
    raise RuntimeError("Not implemented")
    pass


def stats(sol_dir: str):
    dates = list()
    top = list()
    with open(os.path.join(sol_dir, "README.md"), "r") as readme_file:
        lines = readme_file.readlines()
        for i in lines:
            fields = i.split('|')
            if len(fields) < 4 or '?' in fields[2] or '/' not in fields[2]:
                continue
            res = fields[2].split(' / ')
            me, total = int(res[0]), int(res[1])

            top.append(round(me/total*100, 2))
            dates.append(datetime.datetime.strptime(fields[4].title().strip(), DATE_FORMAT.replace("-", "")))

    shared.generate_date_based_plot(os.path.join(sol_dir, ".stats.svg"), dates, top, "Top, %", "Date")


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
        _post_leetcode()
