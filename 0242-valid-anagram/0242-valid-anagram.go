package validanagram

// Compute: O(n)
// Memory: O(n)

//nolint:unused // solution LeetCode problem
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, char := range s {
		m[char]++
	}

	for _, char := range t {
		if m[char] == 0 {
			return false
		}

		m[char]--
	}

	return true
}
