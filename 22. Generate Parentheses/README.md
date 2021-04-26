# Idea

當用完指定數量之左右括號即完成一個組合

透過遞迴的方式走訪各種組合

接著針對這些組合以基於右括號數量不得大於左括號原則下去裁減可行解

將滿足上述限制之解加入res即可


# Time Complexity

需要列舉各種組合應該是 O(N！) ？

# Leetcode Link
[https://leetcode.com/problems/generate-parentheses](https://leetcode.com/problems/generate-parentheses)