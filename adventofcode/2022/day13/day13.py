import os
from functools import cmp_to_key

t_int, t_list = type(0), type(list())


def read_input(filename: str) -> list:
    with open(os.path.join(os.getcwd(), filename), "r") as f:
        return f.readlines()


# 1  - success
# 0  - keep looking
# -1 - failure
def compare(left: list, right: list) -> int:
    l_left, l_right = len(left), len(right)

    for i in range(min(l_left, l_right)):
        tl, tr = type(left[i]), type(right[i])
        x = 0
        if tl == t_int and tr == t_int:
            if left[i] < right[i]:
                x = 1
            elif left[i] > right[i]:
                x = -1
            else:
                x = 0
        elif tl == t_list and tr == t_int:
            x = compare(left[i], [right[i]])
        elif tl == t_int and tr == t_list:
            x = compare([left[i]], right[i])
        elif tl == t_list and tr == t_list:
            x = compare(left[i], right[i])
        if x != 0:
            return x

    if l_left < l_right:
        return 1
    elif l_left > l_right:
        return -1
    else:
        return 0


def day_thirteen_1():
    text = read_input("day13.input")
    pairs = [[eval(text[i]), eval(text[i+1])] for i in range(len(text)) if i == 0 or text[i-1] == "\n"]
    res = 0
    for i, p in enumerate(pairs, start=1):
        if compare(p[0], p[1]) == 1:
            res += i

    print(res)


def day_thirteen_2():
    dp = [[[2]], [[6]]]
    text = read_input("day13.input")
    pairs = [eval(text[i]) for i in range(len(text)) if text[i] != "\n"]
    for i in dp:
        pairs.append(i)
    res = 1
    for i, v in enumerate(sorted(pairs, key=cmp_to_key(compare), reverse=True), start=1):
        if v in dp:
            res *= i
    print(res)


if __name__ == '__main__':
    day_thirteen_1()
    day_thirteen_2()
