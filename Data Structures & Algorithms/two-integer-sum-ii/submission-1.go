func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1


	for left < right {
		currentSum := numbers[left] + numbers[right]

		if currentSum == target {
			return []int{left + 1, right + 1}
		}

		if currentSum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}