# source: https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/
class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vowels = {'a', 'e', 'i', 'o', 'u'}
        cnt = 0
        for i in range(k):
            cnt += int(s[i] in vowels)
        ans = cnt

        for i in range(k, len(s)):
            cnt += int(s[i] in vowels)
            cnt -= int(s[i - k] in vowels)
            ans = max(ans, cnt)

        return ans