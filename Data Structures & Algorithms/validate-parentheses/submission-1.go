func isValid(s string) bool {
    n := len(s)
    if n%2 != 0 {
        return false
    }

    // pre-allocate the exact maximum capacity to guarantee zero slice reallocations.
    stack := make([]byte, 0, n)

    // iterate via index to access raw bytes, avoiding rune decoding.
    for i := 0; i < n; i++ {
        switch s[i] {
        case '(':
            // push the EXPECTED closing bracket instead of the opening one
            stack = append(stack, ')')
        case '{':
            stack = append(stack, '}')
        case '[':
            stack = append(stack, ']')
        default:
            // check for empty stack or mismatched expected closing bracket
            if len(stack) == 0 || stack[len(stack)-1] != s[i] {
                return false
            }
            // pop the top element
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}