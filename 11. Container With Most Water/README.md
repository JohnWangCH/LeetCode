https://leetcode.com/problems/container-with-most-water/description/

# Idea
Using two pointers to find the most capacity.
One is from the highest position and another is from the lowest position.
By moving the pointer with a lower value, we can find the possible combination of max capacity.


# Time Complexity

O(N) - We traverse the input array which contains n elements. 

# Space Complexity

O(1) - No extra space