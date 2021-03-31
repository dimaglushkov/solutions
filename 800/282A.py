n = int(input())
pluses = sum([1 for _ in range(n) if input()[1] == '+'])
print(pluses - (n - pluses))
