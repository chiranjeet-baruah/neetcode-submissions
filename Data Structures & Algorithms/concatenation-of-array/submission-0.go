func getConcatenation(nums []int) []int {
    ans := make([]int, 0, 2*len(nums))

	ans = append(ans,nums...)
	ans = append(ans,nums...)

	return ans
}
