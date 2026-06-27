func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix) // Number of rows
	n := len(matrix[0]) // Number of columns

	// Conceptually flatten the 2D matrix into a 1D array.
	// Initialize pointers to the start and end of this imaginary 1D array.
	low := 0
	high := (m*n)-1

	for low <= high	{
		// Calculate the midpoint. Using low + (high - low) / 2 
		// prevents potential integer overflow compared to (low + high) / 2.
		mid := low + (high-low)/2

		// Translate the 1D 'mid' index back into 2D matrix coordinates:
		// The row is the quotient (how many full rows we've passed).
		// The column is the remainder (how far into the current row we are).
		midValue := matrix[mid/n][mid%n]

		// Binary search
		if midValue == target{
			return true
		} else if midValue < target {
			low = mid + 1 // Target is larger, discard the left half
		} else {
			high = mid - 1 // Target is smaller, discard the right half
		}
	}

	return false
}