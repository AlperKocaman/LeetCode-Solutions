class Solution {
public:
    int minCostClimbingStairs(vector<int>& cost) {
        
    	int vectorSize = cost.size();

		int mins[vectorSize] = {0};
		mins[0] = cost[0];
		mins[1] = cost[1]; 

        for(int i=2; i < vectorSize ; ++i){

        	if(mins[i-2] < mins[i-1])
        		mins[i] = mins[i-2] + cost[i];
        	else
        		mins[i] = mins[i-1] + cost[i];	 
        }

        return mins[vectorSize-2] < mins[vectorSize-1] ? mins[vectorSize-2] : mins[vectorSize-1];
    }
};