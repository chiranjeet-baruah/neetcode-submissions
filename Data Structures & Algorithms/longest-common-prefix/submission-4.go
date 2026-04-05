func longestCommonPrefix(strs []string) string {
     prefix := strs[0]
     for _, s := range strs[1:] {
         for !strings.HasPrefix(s, prefix) {
             prefix = prefix[:len(prefix)-1]
         }
     }
     return prefix
 }