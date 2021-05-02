https://leetcode.com/problems/3sum/

# Idea

The initial idea is the same as two sum problem.
The only difference is that we fixed at current element 'A', and looking for the pair ('B, C') from the rest of the elements. 

So the target value of (B+C) now is varied by (target - A).

To reduce the effort of finding pair (B, C) from the rest of the elements, we sort the input array in advance.
Given that the array is already sorted, we can simply use two pointers (one from the lowest, one from the highest) to find the pair. And HashMap is no longer needed in this case.

Node: please remember to skip the duplicated element.

# Time Complexity

O(N^2) - We traverse the input array which contains n elements, and for each visit we also check the rest of the elements to know if a pair exists. the total time cost at most shall be 1+2+3+4+....+N = O(N^2)

# Space Complexity

O(1) - We don't use any extra space instead of the return list.
