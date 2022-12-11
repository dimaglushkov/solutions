import os


def read_input(filename: str) -> list:
    with open(os.path.join(os.getcwd(), filename), "r") as f:
        return f.readlines()


def new_tail_coords(hx, hy, tx, ty) -> list:
    dx, dy = hx - tx, hy - ty
    if abs(dx) == 1 and abs(dy) == 0 or abs(dx) == 0 and abs(dy) == 1 or abs(dx) == 1 and abs(dy) == 1:
        pass
    elif abs(dx) == 2 and abs(dy) == 2:
        tx += int(dx / 2)
        ty += int(dy / 2)
    elif abs(dx) == 2:
        tx += int(dx / 2)
        if abs(dy) == 1:
            ty += dy
    elif abs(dy) == 2:
        ty += int(dy / 2)
        if abs(dx) == 1:
            tx += dx
    return [tx, ty]


def day_nine_1():
    text = read_input("day09.input")
    hx, hy, tx, ty = 0, 0, 0, 0
    coords = set()

    for line in text:
        direction, distance = line.split(" ")[0], int(line.split(" ")[1])
        for i in range(distance):
            if direction == "U":
                hy += 1
            elif direction == "L":
                hx -= 1
            elif direction == "D":
                hy -= 1
            elif direction == "R":
                hx += 1

            tx, ty = new_tail_coords(hx, hy, tx, ty)
            t = (tx, ty)
            coords.add(t)

    print(len(coords))


def day_nine_2():
    text = read_input("day09.input")
    knots = list()
    coords = set()

    for i in range(10):
        knots.append([0, 0])

    for line in text:
        vals = line.split(" ")
        direction, distance = vals[0], int(vals[1])
        for x in range(distance):
            if direction == "U":
                knots[0][1] += 1
            elif direction == "L":
                knots[0][0] -= 1
            elif direction == "D":
                knots[0][1] -= 1
            elif direction == "R":
                knots[0][0] += 1

            i = 1
            while i < len(knots):
                knots[i][0], knots[i][1] = new_tail_coords(knots[i-1][0], knots[i-1][1], knots[i][0], knots[i][1])
                i += 1

            t = (knots[-1][0], knots[-1][1])
            coords.add(t)

    print(len(coords))


def main():
    day_nine_1()
    day_nine_2()


if __name__ == '__main__':
    main()
