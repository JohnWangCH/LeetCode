class Solution {
    public int maxArea(int[] height) {
        int left = 0, right = height.length-1;
        int maxSize = 0;
        while(left < right){
            int dist = right - left;
            if(height[right] > height[left]){
                maxSize = Math.max(maxSize, height[left]*dist);
                left++;
            }
            else{
                maxSize = Math.max(maxSize, height[right]*dist);
                right--;
            }
        }
        return maxSize;
    }
}