func maxArea(heights []int) int {
	i := 0
	j := len(heights)-1

	maxArea := 0

	for i < j {
		maxAreaInternal := (j-i) * min(heights[i],heights[j])

		if maxAreaInternal > maxArea {
			maxArea = maxAreaInternal
		}

		if heights[i] < heights[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}

func min(a,b int) int{
	if a < b {
		return a
	} else {
		return b
	}
}