func reverse(x int) int {
	res := 0

	for x != 0 {
		digit := x % 10

		if res > (math.MaxInt32-digit)/10 {
			return 0
		}

		if res < (math.MinInt32-digit)/10 {
			return 0
		}

		res = res*10 + digit
		x /= 10
	}

	return res
}
