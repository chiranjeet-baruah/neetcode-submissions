func minEatingSpeed(piles []int, h int) int {
    low, high := 1, 0
    for _, p := range piles {
        if p > high {
            high = p
        }
    }

    for low < high {
        mid := low + (high-low)/2
        if canFinish(piles, h, mid) {
            high = mid
        } else {
            low = mid + 1
        }
    }
    return low
}

func canFinish(piles []int, h, k int) bool {
    hours := 0
    for _, p := range piles {
        hours += (p + k - 1) / k
        if hours > h {
            return false
        }
    }
    return true
}