n, m = [int(i) for i in input().split(' ')]
mid = int(n / 2) if not n % 2 else int(n / 2) + 1
print(2 * m - 1 if m > mid else (m - mid) * 2)
