func topKFrequent(nums []int, k int) []int {
      freqMap := make(map[int]int, len(nums))

      for _, num := range nums {
          freqMap[num]++
      }

      buckets := make([][]int, len(nums)+1)

      for num, freq := range freqMap {
          buckets[freq] = append(buckets[freq], num)
      }

      results := make([]int, 0, k)

      for i := len(buckets) - 1; i > 0; i-- {
          results = append(results, buckets[i]...)
          if len(results) >= k {
              return results[:k]
          }
      }
      return results
  }