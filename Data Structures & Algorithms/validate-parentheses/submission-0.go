func isValid(s string) bool {
    // an odd-length string can never be valid
    if len(s)%2 != 0 {
        return false
    }

    var stack []rune
    
    // map to easily match closing brackets to their opening pairs
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, char := range s {
        if char == '(' || char == '{' || char == '[' {
            // push opening brackets onto the stack
            stack = append(stack, char)
        } else {
            // check if stack is empty before trying to access the top element
            if len(stack) == 0 {
                return false
            }
            
            // pop the top element
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            // check if the popped element matches the current closing bracket
            if top != pairs[char] {
                return false
            }
        }
    }

    // if the stack is empty then all brackets were matched
    return len(stack) == 0
}