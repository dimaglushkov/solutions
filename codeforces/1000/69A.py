n = int(input())
v = [0, 0, 0]
for i in range(n):
    for ji, j in enumerate(input().split(' ')):
        v[ji] += int(j)
print('NO' if any(v) else 'YES')
