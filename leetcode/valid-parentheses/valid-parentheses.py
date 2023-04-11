# source: https://leetcode.com/problems/valid-parentheses/
class Solution:
    def isValid(self, s: str) -> bool:
        pairs = {
            "}": "{",
            "]": "[",
            ")": "("
        }
        stack = list()
        for c in s:
            if c in pairs.values():
                stack.append(c)
            else:
                if len(stack) == 0 or stack.pop() != pairs[c]:
                    return False
        return len(stack) == 0


Solution().isValid("()[]{}")
