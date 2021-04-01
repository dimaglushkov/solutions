s = input()
cs, ci = 'hello', 0
for i in range(len(s)):
    if s[i] == cs[ci]:
        ci += 1
    if ci == 5:
        print('YES')
        exit(0)
print('NO')




# s = [s[i] for i in range(len(s) - 1) if s[i] != s[i - 1]]
# pass
