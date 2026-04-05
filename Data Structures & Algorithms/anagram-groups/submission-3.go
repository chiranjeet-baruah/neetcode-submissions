func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[[26]byte][]string)

    for _, word := range strs {
        var charCount [26]byte
        
        for i := 0; i < len(word); i++ {
            charCount[word[i]-'a']++
        }

        anagramMap[charCount] = append(anagramMap[charCount], word)
    }

    result := make([][]string, 0, len(anagramMap))
    for _, group := range anagramMap {
        result = append(result, group)
    }

    return result
}