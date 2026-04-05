func hasDuplicate(nums []int) bool {
    map1 := make(map[int]struct{})

	for _, value := range(nums) {
		_, found := map1[value]
		if !found {
			map1[value] = struct{}{}
			continue
		} 
		return true
	}
	return false
}
