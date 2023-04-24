# source: https://leetcode.com/problems/last-stone-weight/
import heapq
from typing import List


class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        stones = [-1 * s for s in stones]
        heapq.heapify(stones)
        while len(stones) > 1:
            x, y = heapq.heappop(stones), heapq.heappop(stones)
            x, y = max(x, y), min(x, y)
            if x != y:
                heapq.heappush(stones, y - x)
        if len(stones) == 0:
            return 0
        return -1 * stones[0]


print(Solution().lastStoneWeight([7, 6, 7, 6, 9]))
print(Solution().lastStoneWeight([2, 7, 4, 1, 8, 1]))
