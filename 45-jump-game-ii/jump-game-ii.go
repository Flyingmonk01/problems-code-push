
func helper(nums []int, n, idx int, dp []int) int {
    if idx >= n-1 {
        return 0
    }
    if dp[idx] != -1 {
        return dp[idx]
    }
    mini := 1000000000
    jump := nums[idx]

    for i := idx + 1; i <= min(idx + jump, n-1); i++ {
        val := 1 + helper(nums, n, i, dp)
        mini = min(mini, val)
    }
    dp[idx] = mini
    return dp[idx]
}

func jump(nums []int) int {
    var n int = len(nums)
    dp := make([]int, n+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = -1
    }
    return helper(nums, n, 0, dp)
}
