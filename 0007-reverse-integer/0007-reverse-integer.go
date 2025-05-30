package reverseinteger

// Compute: O(log x)
// Memory: O(1)

//nolint:unused, mnd // solution LeetCode problem
func reverse(x int) int {
	const (
		maxInt32 = 2147483647
		minInt32 = -2147483648
	)

	res := 0

	for x != 0 {
		digit := x % 10

		if res > maxInt32/10 || (res == maxInt32/10 && digit > 7) {
			return 0
		}

		if res < minInt32/10 || (res == minInt32/10 && digit < -8) {
			return 0
		}

		res = res*10 + digit
		x /= 10
	}

	return res
}
