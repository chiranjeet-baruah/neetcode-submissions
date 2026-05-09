func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}

    for i := 0; i < len(nums)-2; i++ {
        // If the smallest number is > 0, sum can't be 0
        if nums[i] > 0 {
            break
        }
        
        // Skip duplicates for the first element
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        j := i + 1
        k := len(nums) - 1

        for j < k {
            sum := nums[i] + nums[j] + nums[k]

            if sum < 0 {
                j++
            } else if sum > 0 {
                k--
            } else {
                // Found a valid triplet!
                res = append(res, []int{nums[i], nums[j], nums[k]})
                
                // Skip duplicates for j
                for j < k && nums[j] == nums[j+1] {
                    j++
                }
                // Skip duplicates for k
                for j < k && nums[k] == nums[k-1] {
                    k--
                }
                
                // Move past the last duplicate we were sitting on
                j++
                k--
            }
        }
    }
    return res
}