func search(nums []int, target int) int {
    low, high := 0, len(nums)-1
    
    for low <= high {
        mid := low + (high-low)/2
        
        if nums[mid] == target {
            return mid
        }
        
        if nums[mid] < target {
            low = mid + 1 // Target is in the right half
        } else {
            high = mid - 1 // Target is in the left half
        }
    }
    
    return -1
}