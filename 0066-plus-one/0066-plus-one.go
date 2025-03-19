package plusone

// Compute: O(n)
// Memory: O(1)

//nolint:unused // solution LeetCode problem
func plusOne(digits []int) []int {
	maxDigit := 9
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < maxDigit {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}
