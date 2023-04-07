# source: https://leetcode.com/problems/number-of-enclaves/
from typing import List


class Solution:
    def numEnclaves(self, grid: List[List[int]]) -> int:
        ans = 0
        n, m = len(grid), len(grid[0])
        visited = [[False for _ in range(m)] for _ in range(n)]
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        def bfs(i: int, j: int):
            queue = [(i, j)]
            while len(queue) > 0:
                x, y = queue.pop(0)
                if visited[x][y]:
                    continue
                visited[x][y] = True
                for dx, dy in directions:
                    nx, ny = x + dx, y + dy
                    if 0 <= nx < n and 0 <= ny < m and grid[nx][ny] == 1:
                        queue.append([nx, ny])

        for i in [0, n - 1]:
            for j in range(m):
                if grid[i][j] == 1:
                    bfs(i, j)

        for j in [0, m - 1]:
            for i in range(n):
                if grid[i][j] == 1:
                    bfs(i, j)

        for i in range(n):
            for j in range(m):
                if grid[i][j] == 1 and not visited[i][j]:
                    ans += 1

        return ans


# Solution().numEnclaves([[0, 0, 0, 0], [1, 0, 1, 0], [0, 1, 1, 0], [0, 0, 0, 0]])
Solution().numEnclaves([[0], [1], [1], [0], [0]])
Solution().numEnclaves([[1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1],
                        [0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 0],
                        [0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0],
                        [1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0],
                        [0, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1],
                        [1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 1, 0, 0],
                        [1, 0, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1],
                        [0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 0],
                        [0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 1, 1],
                        [1, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1],
                        [0, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 1],
                        [1, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 0, 1, 0],
                        [1, 1, 1, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 1]])
Solution().numEnclaves([
    [0, 0, 0, 1, 1, 1, 0, 1, 0, 0],
    [1, 1, 0, 0, 0, 1, 0, 1, 1, 1],
    [0, 0, 0, 1, 1, 1, 0, 1, 0, 0],
    [0, 1, 1, 0, 0, 0, 1, 0, 1, 0],
    [0, 1, 1, 1, 1, 1, 0, 0, 1, 0],
    [0, 0, 1, 0, 1, 1, 1, 1, 0, 1],
    [0, 1, 1, 0, 0, 0, 1, 1, 1, 1],
    [0, 0, 1, 0, 0, 1, 0, 1, 0, 1],
    [1, 0, 1, 0, 1, 1, 0, 0, 0, 0],
    [0, 0, 0, 0, 1, 1, 0, 0, 0, 1]
])
