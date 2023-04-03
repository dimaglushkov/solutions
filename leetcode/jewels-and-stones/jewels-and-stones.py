# source: https://leetcode.com/problems/jewels-and-stones/
class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        ans = 0
        x = set()
        for j in jewels:
            x.add(j)

        for s in stones:
            if s in x:
                ans += 1

        return ans