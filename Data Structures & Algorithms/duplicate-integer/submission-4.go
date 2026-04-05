func hasDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := seen[v]; ok {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}