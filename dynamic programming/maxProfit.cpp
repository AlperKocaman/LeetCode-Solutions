/*
Best Time to Buy and Sell Stock
	->find the maximum profit 
*/

/* Brute Force Approach*/
// O(n^2) time complexity
int maxProfit(vector<int>& prices) {

 	int max = 0;

 	for(int i=prices.size()-1; i>0; i--){

 		for(int j=0; j<i; j++){

 			if(prices[i] - prices[j] > max)
 				max = prices[i] - prices[j];
 		}
 	}
 	return max;
 }

/* DP approach */
// O(n) time complexity 
int maxProfitFast(vector<int>& prices){

 	int minPrice = INT_MAX;
 	int maxProfit = 0;

 	for(int i=0; i<prices.size(); i++){

 		if(prices[i] < minPrice)
 			minPrice = prices[i];
 		if(prices[i] - minPrice > maxProfit)
 			maxProfit = prices[i] - minPrice;
 	}

 	return maxProfit;
 }