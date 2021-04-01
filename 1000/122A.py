lns = [4, 7, 44, 47, 74, 77, 444, 447, 474, 477, 744, 747, 774, 777]
n = int(input())
r = [1 for ln in lns if not n % ln]
print('YES' if len(r) else 'NO')
