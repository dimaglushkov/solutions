# source: https://leetcode.com/problems/unique-paths-iii/
# using pythons solution as prototype

from typing import List


class Solution:
    def uniquePathsIII(self, grid: List[List[int]]) -> int:
        count = 0
        rows = len(grid)
        cols = len(grid[0])

        def find_start():
            for i in range(len(grid)):
                for j in range(len(grid[0])):
                    if grid[i][j] == 1:
                        return i, j

        def count_zeros():
            zeros = 0
            for i in range(len(grid)):
                for j in range(len(grid[0])):
                    if grid[i][j] == 0:
                        zeros += 1
            return zeros

        istart, jstart = find_start()

        def dfs(i, j, z):
            if grid[i][j] == 2:
                if z != -1:
                    return
                nonlocal count
                count += 1
                return

            grid[i][j] = 1

            if i + 1 < rows and (grid[i + 1][j] == 0 or grid[i + 1][j] == 2):
                dfs(i + 1, j, z - 1)

            if j + 1 < cols and (grid[i][j + 1] == 0 or grid[i][j + 1] == 2):
                dfs(i, j + 1, z - 1)

            if i - 1 > -1 and (grid[i - 1][j] == 0 or grid[i - 1][j] == 2):
                dfs(i - 1, j, z - 1)

            if j - 1 > -1 and (grid[i][j - 1] == 0 or grid[i][j - 1] == 2):
                dfs(i, j - 1, z - 1)

            grid[i][j] = 0

        dfs(istart, jstart, count_zeros())
        return count




if __name__ == '__main__':
    sol = Solution()
    print("Expected: 2, Output: ", sol.uniquePathsIII([[1, 0, 0, 0], [0, 0, 0, 0], [0, 0, 2, -1]]))
    print("Expected: 4, Output: ", sol.uniquePathsIII([[1, 0, 0, 0], [0, 0, 0, 0], [0, 0, 0, 2]]))
    print("Expected: 0, Output: ", sol.uniquePathsIII([[0, 1], [2, 0]]))
