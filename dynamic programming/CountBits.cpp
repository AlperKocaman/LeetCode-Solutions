class Solution {
public:
    vector<int> countBits(int num) {
   		
   		if(num == 0){
   			return std::vector<int> {0};
   		}

    	std::vector<int> numVector;

      std::cout<<numVector.size();

      numVector.push_back(0);
      numVector.push_back(1);

   		int next2sPower = 2;
   		int current2sPower = 1;

   		for(int i=2 ; i <= num ; ++i){

   			if(i == next2sPower){
   				numVector.push_back(1);
   				next2sPower *= 2;
   				current2sPower *= 2;
   			}

   			else{
   				numVector.push_back(numVector[i-current2sPower] + numVector[current2sPower]);
   			}

   		}

   		return numVector;
    }
};

/*
*Runtime: 52 ms, faster than 93.75% of C++ online submissions for Counting Bits.
*Memory Usage: 9.9 MB, less than 36.59% of C++ online submissions for Counting Bits.
*/