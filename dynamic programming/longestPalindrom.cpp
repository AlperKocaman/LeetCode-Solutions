class Solution {
public:
	
    string longestPalindrome(string s) {
  		
  		int sSize = s.length();

  		if(sSize == 0)
  			return "";
  		string subStr = s.substr(0,1);
  		int subStrSize = 1;
  		int tempSubStrSize = 0;
  		int i=0,j=0,k=0;

  		for(i = 0; i < sSize ; ++i){

  			for(j = i-1 , k = i+1; j >= 0 && k < sSize ; --j , ++k){

  				if(s[j] == s[k])
  					tempSubStrSize += 2;
  				else
  					break;
  			}

  			if(tempSubStrSize +1 > subStrSize){
  				subStrSize = tempSubStrSize+1;
  				subStr = s.substr(j+1,tempSubStrSize+1);
  			}

  			tempSubStrSize = 0;

  			for(j = i , k = i+1; j >= 0 && k < sSize ; --j , ++k){

  				if(s[j] == s[k])
  					tempSubStrSize += 2;
  				else
  					break;
  			}

  			if(tempSubStrSize > subStrSize){
  				subStrSize = tempSubStrSize;
  				subStr = s.substr(j+1,tempSubStrSize);
  			}

  			tempSubStrSize = 0;
  		}

  		return subStr;      
    }
    
};

/*
*Runtime: 16 ms, faster than 85.68% of C++ online submissions for Longest Palindromic Substring.
*Memory Usage: 16 MB, less than 37.93% of C++ online submissions for Longest Palindromic Substring.
*/