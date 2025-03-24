package plusone

// Compute: O(n)
// Memory: O(n)

//nolint:unused // solution LeetCode problem
func plusOne(digits []int) []int {
	maxDigit := 9
	n := len(digits)
	digits = append(digits, 0)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < maxDigit {
			digits[i]++
			return digits[:n]
		}
		digits[i] = 0
	}
	digits[0] = 1
	return digits
}
