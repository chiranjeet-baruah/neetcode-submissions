func hasDuplicate(nums []int) bool {
    map1 := make(map[int]int)

	for _, value := range(nums) {
		map1[value]++
		if map1[value] > 1 {
			return true
		}
	}
	return false
}
