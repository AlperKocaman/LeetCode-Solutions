class Solution {
public:
    bool isSubsequence(string s, string t) {

    	int sSize = s.size();
    	int tSize = t.size();
    	int sIndex = 0;

    	for(int i=0; i < tSize ; ++i){

    		if(t[i] == s[sIndex]){
    			sIndex++;
    		}

    		if(sIndex >= sSize)
    			return true;
    	}

    	if(sIndex >= sSize)
    			return true;
        return false;
    }
};