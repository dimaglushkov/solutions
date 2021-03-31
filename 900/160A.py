n = int(input())
coins = sorted([int(i) for i in input().split(' ')], reverse=True)
print([i for i in range(1, len(coins) + 1) if sum(coins[:i]) > sum(coins[i:])][0])
