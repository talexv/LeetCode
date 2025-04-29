package stringtointegeratoi

import "strings"

//nolint:unused // solution LeetCode problem
func extractSign(s string) (string, int) {
	switch s[0] {
	case '+':
		return s[1:], 1
	case '-':
		return s[1:], -1
	default:
		return s, 1
	}
}

//nolint:unused, mnd // solution LeetCode problem
func myAtoi(s string) int {
	const (
		maxInt32 = 2147483647
		minInt32 = -2147483648
	)

	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}

	s, sign := extractSign(s)
	res := 0

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			break
		}

		digit := int(ch - '0')

		if res > maxInt32/10 || (res == maxInt32/10 && digit > 7) {
			if sign < 0 {
				return minInt32
			}

			return maxInt32
		}

		res = res*10 + digit
	}

	return res * sign
}
