# Idea

scan s and p and determine the result of a match between the current character from s and the character from p.

this scan can done by looping s and p. if reach the end of both strings while there is still a match then return True, otherwise return False.

before we looping the string, need to handle the empty string case

it also can be solved by DP

For "?" case: if two chars are the same or pattern character is ?, we can use the previous result arr2d[i-1][j-1]

For "*" case: if pattern character is *, we can use the result depend on the arr2d[i][j-1] or arr2d[i-1][j]


# Time Complexity

O(len(s) + len(p))

# Leetcode Link
[https://leetcode.com/problems/wildcard-matching](https://leetcode.com/problems/wildcard-matching)