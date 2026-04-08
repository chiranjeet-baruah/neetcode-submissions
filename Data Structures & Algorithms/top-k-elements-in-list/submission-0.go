func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for i:=0;i<len(nums);i++ {
		freqMap[nums[i]]++
	}

	frequencyBuckets := make([][]int,len(nums)+1)

	for number, frequency := range freqMap{
		frequencyBuckets[frequency] = append(frequencyBuckets[frequency],number)
	}

	var results []int

	for i:=len(frequencyBuckets)-1;i>0;i-- {
		for _,number := range frequencyBuckets[i] {
			results = append(results,number)
			if len(results) == k {
				return results
			}
		}
	}
	return results
}