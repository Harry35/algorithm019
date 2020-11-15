package Week_02

func nthUglyNumber(n int) int {
    dp:= make([]int,n,n)
    dp[0] = 1
    p2, p3, p5 := 0, 0, 0
    for i := 1; i < n; i++ {
        dp[i] = min(min(dp[p2] * 2, dp[p3] * 3), dp[p5] * 5)
        if dp[i] == dp[p2] * 2 {
            p2++
        }
        if dp[i] == dp[p3] * 3 {
            p3++
        }
        if dp[i] == dp[p5] * 5 {
            p5++
        }
    }

    return dp[n-1]
}

func min(a, b int) int {
    if a < b {
        return a
    }

    return b
}