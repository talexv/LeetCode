package reverseinteger

import (
	"math"
	"strconv"
)

// Compute: O(n)
// Memory: O(n)

//nolint:unused // solution LeetCode problem
func reverse(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
		x *= sign
	}

	s := strconv.Itoa(x)
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	res, err := strconv.Atoi(string(r))
	if err != nil {
		return 0
	}

	res *= sign
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
