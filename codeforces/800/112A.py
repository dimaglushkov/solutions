row1 = input().lower()
row2 = input().lower()
print(0 if row1 == row2 else -1 if not row1 > row2 else 1)
