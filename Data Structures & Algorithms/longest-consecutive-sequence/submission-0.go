func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Create a Set using a map with empty structs for O(1) lookups
	numsMap := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	maxConsecutive := 0

	// Iterate through the numbers
	for _, num := range nums {
		// Only start counting if this is the beginning of a sequence
		// (Check if num-1 is missing)
		if _, ok := numsMap[num-1]; !ok {
			currentNum := num
			currentStreak := 1

			// Count upwards until the sequence breaks
			for {
				if _, ok := numsMap[currentNum+1]; ok {
					currentNum++
					currentStreak++
				} else {
					break
				}
			}

			// Update the record
			if currentStreak > maxConsecutive {
				maxConsecutive = currentStreak
			}
		}
	}

	return maxConsecutive
}