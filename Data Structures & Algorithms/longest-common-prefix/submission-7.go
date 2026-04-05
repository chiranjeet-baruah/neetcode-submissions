func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i:=1;i<len(strs);i++{
		s := strs[i]
		j := 0

		minLen := len(prefix)
		if len(s) < minLen {
			minLen = len(s)
		}

		for j < minLen && prefix[j] == s[j]{
			j++
		}

		prefix = prefix[:j]
        if j == 0 {
            return ""
        }
    }
	return prefix
}
