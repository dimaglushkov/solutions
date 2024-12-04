import os
import re
import shutil
import sys

from . import shared

def create(values, sol_dir):
    if len(values) != 1 or not values[0].isdecimal():
        print(f"incorrect value: {values}")
        exit(1)

    day = "day" + ("0" if int(values[0]) < 10 else "") + values[0]
    year = max([int(x) for x in os.listdir(sol_dir) if x.isdecimal()])

    if day in set(os.listdir(os.path.join(sol_dir, str(year)))):
        print(f"{day} already exists")
        exit(1)

    dir_path = str(os.path.join(sol_dir, str(year), day))
    os.mkdir(dir_path)

    open(os.path.join(dir_path, f"{day}.example"), "w").close()
    open(os.path.join(dir_path, f"{day}.input"), "w").close()
    shutil.copy(os.path.join(shared.get_templates_dir(), "adventofcode.go"), os.path.join(dir_path, "main.go"))

    with open(os.path.join(sol_dir, str(year), "README.md"), "a") as f:
        f.write(f"| [Day {values[0]}](https://adventofcode.com/{year}/day/{values[0]}) | [golang](https://github.com/dimaglushkov/solutions/tree/master/adventofcode/{year}/{day}/main.go) |")

    print(f"file://{os.path.join(dir_path, "main.go")}")




def run(values, lang, sol_dir, handle, action):
    if action == 'generate':
        create(values, sol_dir)

