func longestConsecutive(nums []int) int {
     if len(nums) == 0 {
         return 0
     }

     numsMap := make(map[int]struct{}, len(nums)) // pre-sized map avoids rehash growth
     for _, v := range nums {
         numsMap[v] = struct{}{}
     }

     maxConsecutive := 0
     for v := range numsMap { // iterate map keys (unique) — skips duplicate sequence-start evaluations
         if _, ok := numsMap[v-1]; !ok {
             currentStreak := 1
             for {
                 if _, ok := numsMap[v+currentStreak]; !ok { // offnumsMap arithmetic avoids currentNum variable
                     break
                 }
                 currentStreak++
             }
             if currentStreak > maxConsecutive {
                 maxConsecutive = currentStreak
             }
         }
     }
     return maxConsecutive
 }