matrix = [[i for i in input().split()] for j in range(5)]
x, y = 0, 0
for row in matrix:
    if '1' in row:
        x = matrix.index(row)
        y = row.index('1')
print(abs(x - 2) + abs(y - 2))
