import os
from math import prod, lcm, floor


def read_input(filename: str) -> list:
    with open(os.path.join(os.getcwd(), filename), "r") as f:
        return f.readlines()


def parse_data(text):
    data = dict()
    i, cur = 0, 0
    while i < len(text):
        line = text[i].strip()
        if line.startswith("Monkey"):
            cur = int(line.split(" ")[-1].replace(":", ""))
            data[cur] = dict()
            data[cur]["insp"] = 0
        elif line.startswith("Starting"):
            data[cur]["items"] = [int(x) for x in line.split(":")[-1].split(",")]
        elif line.startswith("Operation"):
            oper = line.split(" ")[-1]
            data[cur]["action"] = {
                "op": line.split(" ")[-2] if oper != "old" else "**",
                "operand": int(oper) if oper != "old" else None
            }
        elif line.startswith("Test"):
            data[cur]["throw"] = dict()
            data[cur]["throw"]["div"] = int(line.split(" ")[-1])
            data[cur]["throw"][True] = int(text[i + 1].split(" ")[-1])
            data[cur]["throw"][False] = int(text[i + 2].split(" ")[-1])
            i += 2
        i += 1
    return data


def day_eleven_1():
    data = parse_data(read_input("day11.input"))
    for i in range(20):
        for mi, md in data.items():
            if len(md.items()) == 0:
                continue
            ii = 0
            while len(md["items"]) > 0:
                # inspecting items
                data[mi]["insp"] += 1
                action = data[mi]["action"]
                if action["op"] == "+":
                    data[mi]["items"][ii] = floor((data[mi]["items"][ii] + action["operand"]) / 3)
                elif action["op"] == "*":
                    data[mi]["items"][ii] = floor(data[mi]["items"][ii] * action["operand"] / 3)
                elif action["op"] == "**":
                    data[mi]["items"][ii] = floor(data[mi]["items"][ii] * data[mi]["items"][ii] / 3)

                # throwing items
                target_mi = data[mi]["throw"][data[mi]["items"][ii] % data[mi]["throw"]["div"] == 0]
                data[target_mi]["items"].append(data[mi]["items"][ii])
                data[mi]["items"] = data[mi]["items"][1:]

    print(prod(sorted([insp["insp"] for insp in data.values()])[-2:]))


def day_eleven_2():
    file_name = "day11.input"
    data = parse_data(read_input(file_name))
    div = lcm(*(d["throw"]["div"] for d in data.values()))
    for i in range(10_000):
        for mi, md in data.items():
            if len(md.items()) == 0:
                continue
            ii = 0
            while len(md["items"]) > 0:
                # inspecting items
                data[mi]["insp"] += 1
                action = data[mi]["action"]
                if action["op"] == "+":
                    data[mi]["items"][ii] = floor((data[mi]["items"][ii] + action["operand"]) % div)
                elif action["op"] == "*":
                    data[mi]["items"][ii] = floor(data[mi]["items"][ii] * action["operand"] % div)
                elif action["op"] == "**":
                    data[mi]["items"][ii] = floor(data[mi]["items"][ii] * data[mi]["items"][ii] % div)

                # throwing items
                target_mi = data[mi]["throw"][data[mi]["items"][ii] % data[mi]["throw"]["div"] == 0]
                data[target_mi]["items"].append(data[mi]["items"][ii])
                data[mi]["items"] = data[mi]["items"][1:]
    print(prod(sorted([insp["insp"] for insp in data.values()])[-2:]))


def main():
    day_eleven_1()
    day_eleven_2()


if __name__ == '__main__':
    main()
