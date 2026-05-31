func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    results := make([]int, n)
    
    // Start from the second-to-last day and move backwards. 
    // The last day is always 0, which is handled by Go's default initialization.
    for i := n - 2; i >= 0; i-- {
        j := i + 1
        
        for {
            // Found a warmer day
            if temperatures[j] > temperatures[i] {
                results[i] = j - i
                break
            }
            
            // Reached a day that has no warmer future day
            if results[j] == 0 {
                // Therefore, i also has no warmer future day
                break
            }
            
            // Jump ahead to the next known warmer day for j
            j += results[j]
        }
    }
    
    return results
}