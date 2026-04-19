type Solution struct{}

// Format used per string: <length>#<string> (e.g., ["ab", "c"] -> "2#ab1#c")
func (s *Solution) Encode(strs []string) string {
    var sb strings.Builder
    
    for _, str := range strs {
        // Append length prefix and delimiter
        sb.WriteString(strconv.Itoa(len(str)))
        sb.WriteByte('#')
        
        // Append the actual string payload
        sb.WriteString(str)
    }
    
    return sb.String()
}

// Decode deserializes the string back into the original slice.
// It relies strictly on the parsed length prefix to safely extract strings, 
// ensuring that '#' characters inside the strings do not break the logic.
func (s *Solution) Decode(encoded string) []string {
    var res []string
    i := 0 // i points to the start of the current length prefix

    for i < len(encoded) {
        // Step 1: Scan forward to find the delimiter isolating the length prefix
        j := i
        for encoded[j] != '#' {
            j++
        }

        // Step 2: Parse the extracted length prefix into an integer
        length, _ := strconv.Atoi(encoded[i:j])

        // Step 3: Calculate bounds and extract the exact string payload
        start := j + 1
        end := start + length
        res = append(res, encoded[start:end])

        // Step 4: Advance the main pointer to the start of the next chunk
        i = end
    }
    
    return res
}