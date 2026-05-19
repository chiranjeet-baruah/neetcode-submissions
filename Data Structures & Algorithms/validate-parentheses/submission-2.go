func isValid(s string) bool {
    n := len(s)
    // fail for odd-length strings
    if n%2 != 0 {
        return false
    }

    // a valid string can never have more than n/2 opening brackets.
    // we only allocate exactly half the space.
    limit := n / 2
    stack := make([]byte, limit)
    top := -1 // manual stack pointer

    for i := 0; i < n; i++ {
        switch s[i] {
        case '(':
            top++
            // if we exceed n/2 opening brackets then it's mathematically impossible to close them all.
            if top == limit {
                return false
            }
            stack[top] = ')'
        case '{':
            top++
            if top == limit {
                return false
            }
            stack[top] = '}'
        case '[':
            top++
            if top == limit {
                return false
            }
            stack[top] = ']'
        default:
            // fail if stack is empty or mismatch occurs
            if top == -1 || stack[top] != s[i] {
                return false
            }
            top--
        }
    }

    return top == -1
}