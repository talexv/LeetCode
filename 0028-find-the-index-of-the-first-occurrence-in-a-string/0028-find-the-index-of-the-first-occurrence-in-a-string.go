package findtheindexofthefirstoccurrenceinastring

// Compute: O(n*m)
// Memory: O(1)

//nolint:unused // solution LeetCode problem
func strStr(haystack, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
