def r(x):
    return int(x) if not x % 1 else int(x) + 1


n, m, a = [int(i) for i in input().split(' ')]
print((r(n / a) if a < n else 1) * (r(m / a) if a < m else 1))
