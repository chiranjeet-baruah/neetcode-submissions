func twoSum(nums []int, target int) []int {
	map1 := make(map[int]int)

	response := make([]int, 0, 2)

	for index1, value := range nums{
		index2, exists := map1[target-value]

		if exists {
			response = append(response, index2, index1)
			return response
		} else {
			map1[value] = index1
		}
	} 
	return nil
}
