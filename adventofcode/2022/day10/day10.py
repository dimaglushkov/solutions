import os


def read_input(filename: str) -> list:
    with open(os.path.join(os.getcwd(), filename), "r") as f:
        return f.readlines()


def day_ten_1():
    x = 1
    sc = 20
    ticks = list()
    for line in read_input("day10.input"):
        if "noop" in line:
            ticks.append("-")
        else:
            val = int(line.split(" ")[-1])
            ticks.append("-")
            ticks.append(val)

    res = 0
    for i, c in enumerate(ticks):
        if (i+1) == sc:
            res += x * (i+1)
            sc += 40
        if c != "-":
            x += c
    print(res)


def new_sprite(x):
    n = 40
    res = [ch for ch in "........................................"]

    for i in [-1, 0, 1]:
        if 0 <= x+i < n:
            res[x+i] = "#"

    return res


def day_ten_2():
    x = 1
    sprite = new_sprite(x)
    ticks = list()
    for line in read_input("day10.input"):
        if "noop" in line:
            ticks.append("-")
        else:
            val = int(line.split(" ")[-1])
            ticks.append("-")
            ticks.append(val)

    res = ""
    for i, c in enumerate(ticks):
        if i > 0 and i % 40 == 0:
            res += "\n"

        pos = i % 40
        res += sprite[pos]


        if c != "-":
            x += c
            sprite = new_sprite(x)
    print(res)


def main():
    day_ten_1()
    day_ten_2()


if __name__ == '__main__':
    main()
