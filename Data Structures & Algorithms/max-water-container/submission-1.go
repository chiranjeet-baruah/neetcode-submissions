func maxArea(heights []int) int {
    i := 0
    j := len(heights) - 1
    maxWater := 0

    for i < j {
        // Determine the height of the container (the shorter of the two bars)
        h := heights[i]
        if heights[j] < h {
            h = heights[j]
        }

        // Calculate current area and update max if it's a new record
        currentArea := (j - i) * h
        if currentArea > maxWater {
            maxWater = currentArea
        }

        // Optimization: Move the pointers and skip bars that 
        // are shorter than or equal to the current 'h'.
        // They cannot possibly form a larger container with the current width.
        for i < j && heights[i] <= h {
            i++
        }
        for i < j && heights[j] <= h {
            j--
        }
    }

    return maxWater
}