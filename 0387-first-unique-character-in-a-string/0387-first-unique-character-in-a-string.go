package firstuniquecharacterinastring

// Compute: O(n)
// Memory: O(n)

//nolint:unused // solution LeetCode problem
func firstUniqChar(s string) int {
	m := make(map[rune]int)

	for _, char := range s {
		m[char]++
	}

	for i, char := range s {
		if m[char] == 1 {
			return i
		}
	}

	return -1
}
