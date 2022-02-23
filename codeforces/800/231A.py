solutions = 0
for _ in range(int(input())):
    confidences = [int(conf) for conf in input().split(' ')]
    solutions += 1 if sum(confidences) >= 2 else 0
print(solutions)
