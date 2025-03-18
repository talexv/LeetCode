package reversestring

// Compute: O(n)
// Memory: O(1)

//nolint:unused // solution LeetCode problem
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
