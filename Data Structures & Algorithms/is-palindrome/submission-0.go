func isPalindrome(s string) bool {
	// Initialize pointers at the start and end
	left, right := 0, len(s)-1

	for left < right {
		// 1. Move left pointer past non-alphanumeric characters
		for left < right && !IsAlphanumeric(s[left]) {
			left++
		}
		// 2. Move right pointer past non-alphanumeric characters
		for left < right && !IsAlphanumeric(s[right]) {
			right--
		}

		// 3. Compare the characters case-insensitively
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}

		// 4. Move pointers inward to check the next pair
		left++
		right--
	}

	return true
}

// IsAlphanumeric checks if a byte is a letter or a digit using ASCII ranges
func IsAlphanumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z') ||
		(b >= '0' && b <= '9')
}