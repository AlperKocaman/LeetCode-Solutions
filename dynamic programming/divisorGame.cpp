bool divisorGame(int N){

	return N%2 == 0;

	/*
	Since both players move optimally, if Alice loses for N
	she wins for (N+1) because she can play 1 for the 1st move.
		->lose for 1
		->win  for 2
		->lose for 3 because all evens(only 2) smaller than 3 leads Alice win but the Bob in turn
		-> win for 4 because she loses for 3
		.
		.
		.

		->win  for Evens
		->lose for Odds(odds' divisors are odd as well and odd-odd will be even)
	*/
}

bool divisorGameDP(int N){

	bool dp[N+1] = {false}; // initialize all values to false

	for(int i=2; i<=N ; i++){ // apply tabulation(bottom up approach)

		for(int j=1; j<i; j++){ // find the divisors of number i
			
			// if one of the divisors of that number lead to lose, than Alice
			// will win choosing that divisor and subtract it from the number 
			if((i%j == 0) && !dp[i-j])  
				dp[i] = true;
		}
	}

	return dp[N];
} 