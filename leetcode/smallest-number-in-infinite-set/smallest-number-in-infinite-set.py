# source: https://leetcode.com/problems/smallest-number-in-infinite-set/
import heapq


class SmallestInfiniteSet:

    def __init__(self):
        self.cur = 1
        self.list = list()
        self.set = set()

    def popSmallest(self) -> int:
        if len(self.list) > 0:
            x = heapq.heappop(self.list)
            self.set.remove(x)
            return x

        self.cur += 1
        return self.cur - 1

    def addBack(self, num: int) -> None:
        if num < self.cur and num not in self.set:
            self.set.add(num)
            heapq.heappush(self.list, num)


s = SmallestInfiniteSet()
s.addBack(2)
print(s.popSmallest())
print(s.popSmallest())
print(s.popSmallest())
s.addBack(1)
print(s.popSmallest())
print(s.popSmallest())
print(s.popSmallest())
