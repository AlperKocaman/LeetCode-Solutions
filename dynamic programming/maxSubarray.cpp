class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        
        int numsSize = nums.size();
        int maxSubArrays[numsSize] = {0};
        int max = nums[0];

        for(int i=0 ; i < numsSize ;++i){

        	for(int j=0; j <= i ; ++j){

        		maxSubArrays[j] += nums[i];
        		if(maxSubArrays[j] > max)
        			max = maxSubArrays[j];
        	}
        }
        return max;
    }
};