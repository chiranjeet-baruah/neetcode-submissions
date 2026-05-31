func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    results := make([]int, n)
    
    // Pre-allocate the maximum possible size to avoid runtime memory copying
    stack := make([]int, n)
    top := -1 // Manually track the top of the stack

    for currentIndex, currentTemperature := range temperatures {
        // Evaluate using the manual pointer instead of slice lengths
        for top >= 0 && currentTemperature > temperatures[stack[top]] {
            topIndex := stack[top]
            top-- // "Pop" by simply moving the pointer down
            results[topIndex] = currentIndex - topIndex
        }

        // "Push" by moving the pointer up and assigning
        top++
        stack[top] = currentIndex
    }

    return results
}