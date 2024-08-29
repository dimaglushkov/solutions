# source: https://leetcode.com/problems/most-stones-removed-with-same-row-or-column/
from typing import List


class Solution:
    def removeStones(self, stones: List[List[int]]) -> int:
        n = len(stones)
        ds = dict()
        rows, cols = dict(), dict()

        def union(x, y):
            ds[find(x)] = find(y)

        def find(x):
            if ds[x] != x:
                ds[x] = find(ds[x])
            return ds[x]

        for i, s in enumerate(stones):
            ds[(s[0], s[1])] = (s[0], s[1])

            if s[0] not in rows:
                rows[s[0]] = list()
            rows[s[0]].append(i)

            if s[1] not in cols:
                cols[s[1]] = list()
            cols[s[1]].append(i)

        for k, v in rows.items():
            for c in v:
                union((stones[c][0], stones[c][1]), (stones[v[0]][0], stones[v[0]][1]))

        for k, v in cols.items():
            for c in v:
                union((stones[c][0], stones[c][1]), (stones[v[0]][0], stones[v[0]][1]))

        groups = set()

        for i in stones:
            if find((i[0], i[1])) not in groups:
                groups.add(find((i[0], i[1])))

        return n - len(groups)


sol = Solution()
sol.removeStones(stones=[[0, 1], [1, 2], [1, 3], [3, 3], [2, 3], [0, 2]])
sol.removeStones(stones=[[0, 0], [0, 1], [1, 0], [1, 2], [2, 1], [2, 2]])
sol.removeStones(stones=[[0, 0], [0, 2], [1, 1], [2, 0], [2, 2]])
