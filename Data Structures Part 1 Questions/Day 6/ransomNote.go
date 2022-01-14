// Runtime: 11 ms, faster than 54.59% of Go online submissions for Ransom Note.
// Memory Usage: 4.1 MB, less than 39.91% of Go online submissions for Ransom Note.

func canConstruct(ransomNote string, magazine string) bool {
    if len(magazine) < len(ransomNote) {
        return false
    }
    
    m := make(map[rune]int, 0)
    
    for _, r :=range magazine {
        m[r]++
    }
    
    for _, r :=range ransomNote {
        if v, ok := m[r] ; ok && v>0 {
            m[r]--
        } else {
            return false
        }
    }
    
    return true
}