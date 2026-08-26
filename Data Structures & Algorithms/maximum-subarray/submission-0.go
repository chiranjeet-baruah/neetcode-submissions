func maxSubArray(nums []int) int {
	// if array is empty return 0
	if len(nums) == 0 {
		return 0
	}

	// initialize both current and max sum woth the first element of the array
	currentSum := nums[0]
	maxSum := nums[0]

	// iterate through rest of the array starting from index 1
	for i:=1;i<len(nums);i++{
		// we need to decide whether to add the current element to the ongoing subarray,
		// or start a new subarray from this element
		if nums[i] > currentSum + nums[i] {
			currentSum = nums[i] // start new subaray
		} else {
			currentSum = currentSum + nums[i] // just add to current sum
		}

		// update the overall maximum sum if the current tracking sum is larger
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}