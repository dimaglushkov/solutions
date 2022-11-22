import datetime
import os

import requests

from . import shared

TEMPLATES_DIR = shared.get_templates_dir()
LANG_SPECS = shared.get_lang_specs()
CF_NUM_OF_PROBLEMS = 8
DATE_FORMAT = "%-d %b %Y"

def _pre_codeforces(url: str, lang: str, sol_dir: str):
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
    short_name = f"cf-round-{ext_id}"
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

def _post_codeforces():
    pass

def _pre_leetcode(lang: str, sol_dir: str):
    pass

def _post_leetcode():
    pass


def pre(vals: list, lang: str, sol_dir: str):
    if len(vals) < 1:
        return

    url = vals[0]
    if "codeforces" in url:
        _pre_codeforces(url, lang, sol_dir)
    elif "leetcode" in url:
        _pre_leetcode(lang, sol_dir)


def post(vals: list):
    if len(vals) < 1:
        return

    target = vals[0]
    if target == "cf" or target == "codeforces":
        _post_codeforces()
    elif target == "lc" or target == "leetcode":
        _post_leetcode()
