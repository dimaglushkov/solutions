n = int(input())
m, c = 1, 1
s = [int(i) for i in input().split(' ')]
for i in range(1, len(s)):
    c = c + 1 if s[i - 1] <= s[i] else 1
    if c > m:
        m = c
print(m)
