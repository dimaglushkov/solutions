# source: https://leetcode.com/problems/removing-stars-from-a-string/
class Solution:
    def removeStars(self, s: str) -> str:
        stack = list()
        for c in s:
            if c == "*":
                stack.pop()
            else:
                stack.append(c)
        return "".join(stack)
