class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        arr2d=[[False for i in range(len(p)+1)] for j in range(len(s)+1)]
        
        arr2d[0][0] = True
        
        for j in range(1, len(p) + 1):
            arr2d[0][j] = p[j - 1] == '*' and arr2d[0][j - 1]
        
        for i in range(1, len(s)+1):
            for j in range(1, len(p)+1):
                if p[j - 1] == '*':
                    arr2d[i][j] = arr2d[i - 1][j] or arr2d[i][j - 1]
                elif p[j - 1] == '?' or p[j - 1] == s[i - 1]:
                    arr2d[i][j] = arr2d[i - 1][j - 1]
                    
        return arr2d[len(s)][len(p)]