import "slices"

func groupAnagrams(strs []string) [][]string {
	theMap := make(map[string][]string, len(strs))

	finalSlice := make([][]string, 0, len(strs))

	for _, str := range strs{
		runes := []rune(str)
		slices.Sort(runes) 
		sortedStr := string(runes)

		theMap[sortedStr] = append(theMap[sortedStr], str)
	}

	for _, value := range theMap{
		finalSlice = append(finalSlice,value)
	}

	return finalSlice
}
