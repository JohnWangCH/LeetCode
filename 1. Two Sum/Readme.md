https://leetcode.com/problems/two-sum/

# Idea

The basic idea is to traverse the input array and find out the pair.

During traversing the input array, we want to find out if the current element's complement has already shown up before.

Therefore, we use a hashMap to keep each element and its index.

While we iterate the element, we can check the hashMap to know if a complement already exists.

# Time Complexity

O(N) - We traverse the input array which contains n elements. 

And each loop up in HashMap cost O(1) time 

# Space Complexity

O(N) - We use a hashMap to store the element and its index. It stores at most n elements.