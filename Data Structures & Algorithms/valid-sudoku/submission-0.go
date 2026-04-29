func isValidSudoku(board [][]byte) bool {
    // Initialize three tracking arrays for rows, columns, and 3x3 boxes
    // We use [10] to map digits 1-9 directly to indices 1-9
    var rows [9][10]bool
    var cols [9][10]bool
    var boxes [9][10]bool

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            char := board[r][c]
            
            // 1. Skip empty cells
            if char == '.' {
                continue
            }

            // 2. Convert byte character to integer digit
            num := int(char - '0')
            
            // 3. Calculate which of the 9 boxes this cell belongs to
            boxIdx := (r / 3) * 3 + (c / 3)

            // 4. Check for duplicates in the Row, Column, or Box
            if rows[r][num] || cols[c][num] || boxes[boxIdx][num] {
                return false
            }

            // 5. "File the claim": Mark this number as seen for this Row, Col, and Box
            rows[r][num] = true
            cols[c][num] = true
            boxes[boxIdx][num] = true
        }
    }

    // If we finished the loops without returning false, the board is valid
    return true
}