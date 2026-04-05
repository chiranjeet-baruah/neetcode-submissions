func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	a, b := strs[0], strs[len(strs)-1]
	i := 0
	for i < len(a) && i < len(b) && a[i] == b[i] {
		i++
	}
	return a[:i]
}