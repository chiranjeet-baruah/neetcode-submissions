func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)

	// Store final results
	results := make([]int,n)

	// Waiting list of indices from original temperatures list
	stack := []int{}

	for currentIndex, currentTemperature := range temperatures{
		// While waitlist is not empty 
		// and current temperature is greater than most recent waiting day in waitlist
		for len(stack) > 0 && currentTemperature > temperatures[stack[len(stack)-1]] {
			// Get the index of the most recent waiting day
			topIndex := stack[len(stack)-1]

			// Remove it from the stack
			stack = stack[:len(stack)-1]

			// Calcualte the waiting days
			results[topIndex] = currentIndex - topIndex
		}

		// Add current day's index to the waitlist
		stack = append(stack, currentIndex)
	}

	return results
}