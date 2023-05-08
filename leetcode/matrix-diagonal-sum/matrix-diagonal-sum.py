# source: https://leetcode.com/problems/matrix-diagonal-sum/
class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        n = len(mat)
        return sum([mat[i][i] + mat[i][n-1-i] if i != n-i-1 else mat[i][i] for i in range(n)])


