class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []

        def helper(s: str, l: int, r: int):
            if l + r == n * 2:
                res.append(s)
                return

            if l < n:
                helper(s + "(", l + 1, r)
            if r < l:
                helper(s + ")", l, r + 1)

        helper("", 0, 0)
        return res