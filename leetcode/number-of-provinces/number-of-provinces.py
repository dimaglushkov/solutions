# source: https://leetcode.com/problems/number-of-provinces/
from typing import List


class Solution:
    def findCircleNum(self, isConnected: List[List[int]]) -> int:
        n = len(isConnected)
        parent = [i for i in range(n)]
        def find(x: int) -> int:
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]

        def union(x: int, y: int):
            parent[find(x)] = find(y)

        for i in range(n):
            for j, c in enumerate(isConnected[i]):
                if c:
                    union(i, j)

        groups = set()
        for i in range(n):
            groups.add(find(i))
        return len(groups)