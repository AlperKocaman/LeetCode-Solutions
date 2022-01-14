// Runtime: 33 ms, faster than 52.32% of Go online submissions for First Unique Character in a String.
// Memory Usage: 7.8 MB, less than 5.75% of Go online submissions for First Unique Character in a String.

func firstUniqChar(s string) int {
    charsMap := make(map[rune]int, 0)
    charsSlc := make([]rune, 0)
    
    for i, r := range s {
        val, ok := charsMap[r]
        if !ok {
            charsMap[r] = i
            charsSlc = append(charsSlc, r)
        } else if val == -2{
            charsSlc = append(charsSlc, 0)
        } else {
            charsMap[r] = -2 // seen before, -2 is used for marking non unique runes
            charsSlc[val] = 0
            charsSlc = append(charsSlc, 0)0 0 118 0 0 0 0 116 99 0 100 0
        }
    }
    
    for i, v :=range charsSlc {
        if v != 0 {
            return i
        } 
    }
    
    return -1;
}