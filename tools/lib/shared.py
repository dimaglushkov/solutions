import datetime
import json
import os

import matplotlib.pyplot as plt
import matplotlib.dates as mdates

BG_COLOR = "#0f0f23"
FG_COLOR = "#cccccc"


def get_lang_specs():
    with open(os.path.join(os.path.dirname(__file__), "lang_specs.json")) as json_file:
        x = json.load(json_file)
    return x


def get_templates_dir() -> str:
    return os.path.join(os.path.dirname(__file__), "..", "templates")


def generate_date_based_plot(path: str, xs: list, ys: list):
    f = plt.figure()
    f.set_figwidth(9)
    f.set_figheight(5)
    f.patch.set_facecolor(BG_COLOR)

    plt.gca().xaxis.set_major_formatter(mdates.DateFormatter('%d-%m-%y'))
    plt.gca().xaxis.set_major_locator(mdates.DayLocator(interval=15))
    ax = plt.gca()
    ax.set_facecolor(BG_COLOR)
    for i in ["bottom", "left"]:
        ax.spines[i].set_color(FG_COLOR)
    for i in ["top", "right"]:
        ax.spines[i].set_color(BG_COLOR)
    ax.tick_params(axis='x', colors=FG_COLOR)
    ax.tick_params(axis='y', colors=FG_COLOR)
    plt.plot(xs, ys)
    plt.scatter(xs, ys)
    plt.gcf().autofmt_xdate()
    plt.savefig(path)


def generate_pie_chart(path: str, data: dict):
    fields = list()
    values = list()
    for k, v in data.items():
        fields.append(k)
        values.append(v)

    fig1, ax1 = plt.subplots()
    fig1.patch.set_facecolor(BG_COLOR)
    plt.rcParams['text.color'] = FG_COLOR
    _, _, autotexts = ax1.pie(values, labels=fields, autopct='%1.1f%%', startangle=90)
    for autotext in autotexts:
        autotext.set_color('white')
    ax1.axis('equal')
    my_circle = plt.Circle((0, 0), 0.3, color=BG_COLOR)
    p = plt.gcf()
    p.gca().add_artist(my_circle)
    plt.savefig(path)


def generate_svg_stats(data: dict, sol_dir: str, charts: list) -> str:
    stats = {
        'solved': 0,
        'by_difficulty': dict(),
        'by_tags': dict(),
        'by_lang': dict()
    }

    for problem, meta in data.items():
        stats['solved'] += 1
        if meta['difficulty'] not in stats['by_difficulty']:
            stats['by_difficulty'][meta['difficulty']] = 0
        stats['by_difficulty'][meta['difficulty']] += 1

        for tag in meta['tags']:
            if tag not in stats['by_tags']:
                stats['by_tags'][tag] = 0
            stats['by_tags'][tag] += 1

        for lang in meta['lang']:
            if lang not in stats['by_lang']:
                stats['by_lang'][lang] = 0
            stats['by_lang'][lang] += 1

    stats['by_tags']['other'] = 0
    to_del = list()
    for tag, count in stats['by_tags'].items():
        if count / stats['solved'] < 0.05:
            stats['by_tags']['other'] += count
            to_del.append(tag)

    for tag in to_del:
        del stats['by_tags'][tag]

    target_dir = sol_dir.split("/")[-1]
    headers = list()
    seps = list()
    links = list()
    for c in charts:
        headers.append(f'Solutions {c.replace("_", " ")}')
        seps.append("-")
        links.append(f"![{c}](https://github.com/dimaglushkov/solutions/blob/master/{target_dir}/.{c}.svg)")
        generate_pie_chart(os.path.join(sol_dir, f'.{c}.svg'), stats[c])

    return f'''

Problems solved in total: {stats["solved"]}

|{"|".join(headers)}|
|{"|".join(seps)}|
|{"|".join(links)}|

'''
