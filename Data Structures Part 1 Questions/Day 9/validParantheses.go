// Runtime: 4 ms, faster than 12.62% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.1 MB, less than 84.03% of Go online submissions for Valid Parentheses.

func isValid(s string) bool {
    
    stack := []rune{}
    
    for _,r := range s {
        if isLeftBracket(r) {
            stack = append(stack, r)
        } else {
            if len(stack) < 1 {
                return false
            }
            
            if !isSameType(stack[len(stack)-1], r) {
                return false
            }
            stack = stack[:len(stack) -1]
        }
    }
    
    return len(stack) == 0
}

func isLeftBracket(r rune) bool {
    if r == '{' || r == '(' || r == '[' {
        return true
    }
    
    return false
}

func isSameType(lastInStack, current rune) bool {
    if lastInStack == '{' && current == '}' {
       return true 
    } else if lastInStack == '(' && current == ')'{
        return true 
    } else if lastInStack == '[' && current == ']'{
        return true 
    }
    
    return false
}
