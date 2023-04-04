# source: https://leetcode.com/problems/optimal-partition-of-string/
class Solution:
    def partitionString(self, s: str) -> int:
        ans = 0
        j = 0
        for i, l in enumerate(s):
            if l in s[j:i]:
                ans += 1
                j = i
        return ans + 1
