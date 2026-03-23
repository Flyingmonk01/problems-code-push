func maxProductPath(grid [][]int) int {
    m, n := len(grid), len(grid[0]);

    minP := make([][]int, m)
    maxP := make([][]int, m)

    for i := range minP {
        minP[i] = make([]int, n)
        maxP[i] = make([]int, n)
    }

    minP[0][0] = grid[0][0]
    maxP[0][0] = grid[0][0]

    for j := 1; j < n; j++ {
        minP[0][j] = minP[0][j-1] * grid[0][j]
        maxP[0][j] = maxP[0][j-1] * grid[0][j]
    }

    for i := 1; i < m; i++ {
        minP[i][0] = minP[i-1][0] * grid[i][0]
        maxP[i][0] = maxP[i-1][0] * grid[i][0]
    }

    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            a := minP[i][j-1] * grid[i][j]
            b := maxP[i][j-1] * grid[i][j]
            c := minP[i-1][j] * grid[i][j]
            d := maxP[i-1][j] * grid[i][j]

            minP[i][j] = min(min(a, b), min(c, d))
            maxP[i][j] = max(max(a, b), max(c, d))
        }
    }

    res := maxP[m-1][n-1]

    if res < 0 {
        return -1
    }

    return res % 1000000007
}