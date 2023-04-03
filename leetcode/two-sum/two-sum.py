# source: https://leetcode.com/problems/two-sum/
from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        d = dict()
        for i, v in enumerate(nums):
            if v in d:
                return [d[v], i]
            d[target - v] = i
