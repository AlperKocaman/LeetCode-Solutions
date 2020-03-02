class Solution {
public:
    int rob(vector<int>& nums) {
        
        int numsSize = nums.size();

        if(numsSize == 0)
        	return 0;

        int maxProfit[numsSize] = {0};

        if(numsSize > 2){
        	maxProfit[0] = nums[0];
        	maxProfit[1] = nums[1];
        	maxProfit[2] = nums[2]+ maxProfit[0];
        }

		else if(numsSize == 2){
        	maxProfit[0] = nums[0];
        	maxProfit[1] = nums[1];
        }

        else if(numsSize ==1){
        	return nums[0];	
        }

        for(int i=3 ; i < numsSize ; ++i){

    		if(maxProfit[i-2] > maxProfit[i-3])
    			maxProfit[i] = maxProfit[i-2] + nums[i];
    		else
    			maxProfit[i] = maxProfit[i-3] + nums[i];
		}

		return maxProfit[numsSize-1] > maxProfit[numsSize-2] ? maxProfit[numsSize-1]:maxProfit[numsSize-2];
    }
};

/*
*Runtime: 0 ms, faster than 100.00% of C++ online submissions for House Robber.
*Memory Usage: 8.4 MB, less than 100.00% of C++ online submissions for House Robber.
*/