class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        tail = m + n - 1
        while n - 1 >= 0:
            val = 0
            if m - 1 >= 0 and nums1[m - 1] > nums2[n - 1]:
                val = nums1[m - 1]
                m -= 1
            else:
                val = nums2[n - 1]
                n -= 1

            nums1[tail] = val
            tail -= 1