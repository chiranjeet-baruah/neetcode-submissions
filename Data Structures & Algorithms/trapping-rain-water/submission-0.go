func trap(height []int) int {
    n := len(height)
    if n == 0 {
        return 0
    }

    leftMax := make([]int, n)
    rightMax := make([]int, n)

    // Calculate leftMax
    leftMax[0] = height[0]
    for i := 1; i < n; i++ {
        leftMax[i] = max(leftMax[i-1], height[i])
    }

    // Fill rightMax
    rightMax[n-1] = height[n-1]
    for i := n - 2; i >= 0; i-- {
        rightMax[i] = max(rightMax[i+1], height[i])
    }

    // Calculate total water
    totalWater := 0
    for i := 0; i < n; i++ {
        // Since leftMax and rightMax include height[i], 
        // the result will never be negative.
        totalWater += min(leftMax[i], rightMax[i]) - height[i]
    }

    return totalWater
}

func max(a, b int) int { if a > b { return a }; return b }
func min(a, b int) int { if a < b { return a }; return b }