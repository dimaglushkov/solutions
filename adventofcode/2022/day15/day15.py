import os


def read_input(filename: str) -> list:
    with open(os.path.join(os.getcwd(), filename), "r") as f:
        return f.readlines()


def parse_input(lines: list) -> list:
    sensors = list()
    for line in lines:
        words = line.split(":")[0].split(" ")
        sx, sy = [int(w.split("=")[-1].replace(",", "")) for w in words if "=" in w]
        words = line.split(":")[-1].split(" ")
        bx, by = [int(w.split("=")[-1].replace(",", "")) for w in words if "=" in w]
        sensors.append([sx, sy, bx, by])
    return sorted(sensors)


def distance(x1, y1, x2, y2):
    return abs(x1 - x2) + abs(y1 - y2)


def day_fifteen_1():
    sensors = parse_input(read_input("day15.input"))
    res, y = 0, 2_000_000
    for x in range(-1_000_000, 6_000_000):
        for sx, sy, bx, by in sensors:
            if (bx != x or by != y) and distance(sx, sy, bx, by) >= distance(sx, sy, x, y):
                res += 1
                break
    print(res)


def day_fifteen_2():
    sensors = parse_input(read_input("day15.input"))
    x_cord = 4_000_000
    for y in range(x_cord+1):
        x = 0
        for sx, sy, bx, by in sensors:
            if distance(sx, sy, bx, by) >= distance(sx, sy, x, y):
                x = sx + distance(sx, sy, bx, by) - abs(sy - y) + 1
        if x <= x_cord:
            print(x * x_cord + y)


if __name__ == '__main__':
    day_fifteen_1()
    day_fifteen_2()
