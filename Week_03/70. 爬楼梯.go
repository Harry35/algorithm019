package Week_03

func climbStairs(n int) int {
    one, two := 1, 2
	if n < 3 {
		return n
	}
	for i := 3; i <= n; i++ {
		one, two = two, one + two
	}
	return two
}
