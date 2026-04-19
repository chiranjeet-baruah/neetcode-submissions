func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        // Initialize product to 1
        prod := 1
        for j := 0; j < len(nums); j++ {
            if i != j {
                prod = prod * nums[j]
            }
        }
        res[i] = prod
    }
    return res
}