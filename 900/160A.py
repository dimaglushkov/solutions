n = int(input())
coins = sorted([int(i) for i in input().split(' ')], reverse=True)
print([i for i, coin in enumerate(coins, 1) if sum(coins[:i]) > sum(coins[i:])][0])
