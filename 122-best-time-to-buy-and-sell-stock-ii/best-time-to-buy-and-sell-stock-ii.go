func maxProfitHelper(prices []int, idx, isSell int, dp [][]int) int {
    if idx >= len(prices) {
        return 0;
    }
    if dp[idx][isSell] != -1 {
        return dp[idx][isSell]
    }
    maxi := 100000000

    if isSell == 1 {
        curr_val := prices[idx] + maxProfitHelper(prices, idx+1, 0, dp)
        not_curr_val := maxProfitHelper(prices, idx+1, 1, dp)
        maxi = max(curr_val, not_curr_val)
    }else {
        curr_val := -prices[idx] + maxProfitHelper(prices, idx+1, 1, dp)
        not_curr_val := maxProfitHelper(prices, idx+1, 0, dp)
        maxi = max(curr_val, not_curr_val)
    }
    dp[idx][isSell] = maxi
    return maxi
}

func maxProfit(prices []int) int {
    idx, isSell := 0, 0
    n := len(prices)
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, 2)
        dp[i][0], dp[i][1] = -1, -1
    }
    
    return maxProfitHelper(prices, idx, isSell, dp)
}