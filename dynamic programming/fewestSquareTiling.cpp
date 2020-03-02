class Solution {
public:

	int tilingRectangle(int n, int m){
		int temp;
        if(n>m){
            temp = n;
            n = m;
            m = temp;
        }	
			int **mins = new int *[n];
			for(int i = 0; i <n; i++)
    			mins[i] = new int[m];
    		for(int i = 0 ; i < n ; ++i){
    			for(int j=0; j < m ; ++j){
    				mins[i][j] = 169;
    			}
    		}
			return tilingRectangleHelper(n,m,mins);
	}
    int tilingRectangleHelper(int n, int m, int **mins) {
        
        int temp;
        if(n>m){
            temp = n;
            n = m;
            m = temp;
        }

        if(n == 0)
        	return 0;
        if(n == 1){
        	mins[0][m-1] = m;
        	return m;
        }
        
        for(int i = 1 ; i <= n ; i++){

    		mins[n-1][m-1] = std::min(mins[n-1][m-1],tilingRectangleHelper(n-i,i,mins) + tilingRectangleHelper(i, m-i,mins) + tilingRectangleHelper(n-i, m-i,mins) + 1);
    		return mins[n-1][m-1];
        	
        }
        
        return mins[n][m];
    }
};