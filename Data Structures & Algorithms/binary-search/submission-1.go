func search(nums []int, target int) int {
    return recursiveBinarySearch(0, len(nums)-1, nums, target)
}

func recursiveBinarySearch(low int, high int, nums []int, target int) int {
    if low > high {
        return -1
    }
    
    // Calculate mid safely to prevent integer overflow
    mid := low + (high-low)/2
    
    if nums[mid] == target {
        return mid
    }
    
    if target < nums[mid] {
        return recursiveBinarySearch(low, mid-1, nums, target)
    } 
    
    return recursiveBinarySearch(mid+1, high, nums, target)
}