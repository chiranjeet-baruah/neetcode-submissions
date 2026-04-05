func longestCommonPrefix(strs []string) string {
     minLen := len(strs[0])
     for _, s := range strs[1:] {
         if len(s) < minLen {
             minLen = len(s)
         }
     }
     lo, hi := 0, minLen
     for lo < hi {
         mid := (lo + hi + 1) >> 1
         p := strs[0][:mid]
         ok := true
         for _, s := range strs[1:] {
             if s[:mid] != p {
                 ok = false
                 break
             }
         }
         if ok {
             lo = mid
         } else {
             hi = mid - 1
         }
     }
     return strs[0][:lo]
 }