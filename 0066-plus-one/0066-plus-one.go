package plusone

// Compute: O(n)
// Memory: O(n)

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
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}
