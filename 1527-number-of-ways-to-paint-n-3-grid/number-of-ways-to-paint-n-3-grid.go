const MOD int = 1_000_000_007
var states []string = []string{"RYR", "RYG", "RGR", "RGY", "YRY", "YRG", "YGR", "YGY", "GRY", "GRG", "GYR", "GYG"}
var dp [][]int

func solve(n int, prev int) int {
    if n == 0 {
        return 1
    }

    if dp[n][prev] != -1 {
        return dp[n][prev]
    }

    var result int = 0
    lastPat := states[prev]

    for curr := 0; curr < 12; curr++ {
        if curr == prev {
            continue
        }

        currPat := states[curr]
        conflict := false

        for col := 0; col < 3; col++ {
            if currPat[col] == lastPat[col]{
                conflict = true
                break
            }
        }

        if !conflict {
            result = (result + solve(n-1, curr)) % MOD
        }
    }
    dp[n][prev] = result
    return dp[n][prev]
}

func numOfWays(n int) int {
    var result int = 0

    if n == 1 {
        return 12
    }
    dp = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 12)
		for j := 0; j < 12; j++ {
			dp[i][j] = -1
		}
	}

    for i := 0; i < 12; i++ {
        result = (result + solve(n-1, i)) % MOD
    }
    return result
}