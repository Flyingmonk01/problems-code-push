
func minimumDeleteSumHelper(s1 string, s2 string, i int, j int, dp [][]int) int {
    if i >= len(s1) && j >= len(s2) {
        return 0
    }
    if dp[i][j] != -1 {
        return dp[i][j]
    }

    if i == len(s1){
        d := 0
        for idx := j; idx < len(s2); idx++{
            d += int(s2[idx])
        }
        return d
    }
    if j == len(s2){
        d := 0
        for idx := i; idx < len(s1); idx++{
            d += int(s1[idx])
        }
        return d
    }

    ans := math.MaxInt
    if s1[i] == s2[j] {
        // include
        ans = minimumDeleteSumHelper(s1, s2, i+1, j+1, dp)
    }else{
        // exclude
        ans = min(ans, int(s1[i]) + minimumDeleteSumHelper(s1, s2, i+1, j, dp))
        ans = min(ans, int(s2[j]) + minimumDeleteSumHelper(s1, s2, i, j+1, dp))
    }

    // return ans
    dp[i][j] = ans
    return dp[i][j]

}

func minimumDeleteSum(s1 string, s2 string) int {
    dp := make([][]int, len(s1)+1)
    for i := 0; i <= len(s1); i++ {
        dp[i] = make([]int, len(s2)+1)
        for j := 0; j <= len(s2); j++ {
            dp[i][j] = -1
        }
    }
    return minimumDeleteSumHelper(s1, s2, 0, 0, dp)
}