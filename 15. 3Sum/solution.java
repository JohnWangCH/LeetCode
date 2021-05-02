public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> ret = new ArrayList<>();
        for(int i = 0; i < nums.length; i++){
            if(i == 0 || nums[i] != nums[i-1]){ //skip掉重複值的i
                int target = 0 - nums[i];
                int lo = i+1, hi = nums.length - 1;
                while(lo < hi){
                    if(nums[lo] + nums[hi] == target){
                        ret.add(Arrays.asList(nums[i], nums[lo], nums[hi]));
                        while(lo < hi && nums[lo] == nums[lo+1]) lo++; //skip掉重複值的lo
                        while(lo < hi && nums[hi] == nums[hi-1]) hi--; //skip掉重複值的hi
                        lo++;
                        hi--;
                    }
                    else if (nums[lo] + nums[hi] > target){
                        hi--;
                    }
                    else {
                        lo++;
                    }
                }
            }
        }
        return ret;
    }