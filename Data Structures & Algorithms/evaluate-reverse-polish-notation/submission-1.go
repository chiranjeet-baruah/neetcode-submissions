func evalRPN(tokens []string) int {
	// Pre-allocate the exact maximum capacity needed for the stack
	// to avoid dynamic resizing overhead.
	capacity := (len(tokens) / 2) + 1
	stack := make([]int, 0, capacity)

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			// Pop the top two elements
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			// Evaluate and push the result
			switch token {
			case "+":
				stack = append(stack, left+right)
			case "-":
				stack = append(stack, left-right)
			case "*":
				stack = append(stack, left*right)
			case "/":
				stack = append(stack, left/right)
			}
		default:
			// Push numbers to the stack
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}

	return stack[0]
}