func insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i := 0
	
	for i<len(intervals) && intervals[i][1] < newInterval[0]{
		// append to result
		result = append(result, intervals[i])
		i++
	}

	// i pointer has stopped exactly at the first interval that might overlap with newInterval
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
    	newInterval[0] = min(intervals[i][0], newInterval[0])
    	newInterval[1] = max(intervals[i][1], newInterval[1])
    	i++
	}
	result = append(result, newInterval)

	// append all remaining intervals
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}