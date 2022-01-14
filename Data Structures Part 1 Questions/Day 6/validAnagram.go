// Runtime: 8 ms, faster than 56% of Go online submissions for Valid Anagram.
// Memory Usage: 2.8 MB, less than 52% of Go online submissions for Valid Anagram.

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    m := make(map[rune]int, 0)
    
    for _, r := range s {
        if v, ok := m[r]; ok {
            m[r] = v+1
        } else {
            m[r] = 1
        }
    }
    
    for _, r2 := range t {
        if v, ok := m[r2]; ok && v > 0{
            m[r2] = v-1
        } else {
            return false
        }
    }
    
    return true
}