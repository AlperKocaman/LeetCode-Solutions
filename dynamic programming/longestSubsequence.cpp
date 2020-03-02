class Solution {
public:

    int longestCommonSubsequence(string text1, string text2){
        unsigned int text1length = text1.length();
        unsigned int text2length = text2.length();

        if(!text1length || !text2length)
            return 0;

        int i=0,j=0;
        int **sequences = new int*[text1length];
        for(i = 0; i < text1length ; ++i){
            sequences[i] = new int[text2length];
        }
        for(i = 0; i < text1length ; ++i){
            for(j = 0; j < text2length ; ++j){
                sequences[i][j] = -1;
            }
        }
        return longestCommonSubsequenceHelper(text1, text2, sequences);

    }
    int longestCommonSubsequenceHelper(string text1, string text2, int **sequences) {
        unsigned int text1length = text1.length();
        unsigned int text2length = text2.length();
        
        if(!text1length || !text2length)
            return 0;
        else if(sequences[text1length-1][text2length-1] != -1)
            return sequences[text1length-1][text2length-1];
        else if(text1[text1length-1] == text2[text2length-1])
            return longestCommonSubsequenceHelper(text1.substr(0,text1length-1), text2.substr(0,text2length-1), sequences) +1 ;
        else
            return sequences[text1length-1][text2length-1] = std::max(longestCommonSubsequenceHelper(text1.substr(0,text1length-1), text2, sequences)
                ,longestCommonSubsequenceHelper(text1,text2.substr(0,text2length-1), sequences)) ;
        
    }
};

/*
* Memory limit exceeded!
*/