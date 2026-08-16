func majorityElement(nums []int) int {
	elementMap := make(map[int]int)
	totalNums := len(nums)
	countCheck := totalNums/2

	for _,num := range nums{
		elementMap[num]++
			if elementMap[num] > countCheck {
				return num
			}
	}
	return 0
}