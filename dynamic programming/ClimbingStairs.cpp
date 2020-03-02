class Solution {
public:
    int climbStairs(int n) {

    	int ways[n+1] = {0};
    	ways[1] = 1;
    	ways[0] = 0;
  		return climbStairsHelper(n, ways);	      
    }

    int climbStairsHelper(int n, int ways[]){

    	if(ways[n] != 0)
    		return ways[n];
    	else if(n == 0)
    		return 0;
    	else 
    		ways[n] = climbStairsHelper(n-1,ways) + climbStairsHelper(n-2, ways);
    	return ways[n];	
    }
};