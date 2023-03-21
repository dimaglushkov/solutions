class Solution:
    def zeroFilledSubarray(self, nums: list[int]) -> int:
        ans, c = 0, 0
        for i in nums:
            if i:
                c = 0
            else:
                c += 1
            ans += c
        return ans
