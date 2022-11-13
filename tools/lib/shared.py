import os

from matplotlib import pyplot as plt


def generate_chart(path: str, data: dict):
    fields = list()
    values = list()
    for k, v in data.items():
        fields.append(k)
        values.append(v)

    fig1, ax1 = plt.subplots()
    fig1.patch.set_facecolor('gray')
    plt.rcParams['text.color'] = 'white'
    ax1.pie(values, labels=fields, autopct='%1.1f%%', startangle=90)
    ax1.axis('equal')
    my_circle = plt.Circle((0, 0), 0.3, color='gray')
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
        generate_chart(os.path.join(sol_dir, f'.{c}.svg'), stats[c])

    return f'''

Problems solved in total: {stats["solved"]}

|{"|".join(headers)}|
|{"|".join(seps)}|
|{"|".join(links)}|

'''
