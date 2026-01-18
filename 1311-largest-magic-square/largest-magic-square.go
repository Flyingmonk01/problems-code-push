func largestMagicSquare(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])

    // Prefix sums
    rowSum := make([][]int, m)
    colSum := make([][]int, m+1)

    for i := 0; i < m; i++ {
        rowSum[i] = make([]int, n+1)
    }
    for i := 0; i <= m; i++ {
        colSum[i] = make([]int, n)
    }

    // Build prefix sums
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            rowSum[i][j+1] = rowSum[i][j] + grid[i][j]
            colSum[i+1][j] = colSum[i][j] + grid[i][j]
        }
    }

    maxK := min(m, n)

    // Try largest square first
    for k := maxK; k >= 2; k-- {
        for r := 0; r+k <= m; r++ {
            for c := 0; c+k <= n; c++ {
                if isMagic(grid, rowSum, colSum, r, c, k) {
                    return k
                }
            }
        }
    }

    return 1
}

func isMagic(grid [][]int, rowSum, colSum [][]int, r, c, k int) bool {
    // Target sum = first row
    target := rowSum[r][c+k] - rowSum[r][c]

    // Rows
    for i := 0; i < k; i++ {
        if rowSum[r+i][c+k]-rowSum[r+i][c] != target {
            return false
        }
    }

    // Columns
    for j := 0; j < k; j++ {
        if colSum[r+k][c+j]-colSum[r][c+j] != target {
            return false
        }
    }

    // Diagonals
    d1, d2 := 0, 0
    for i := 0; i < k; i++ {
        d1 += grid[r+i][c+i]
        d2 += grid[r+i][c+k-1-i]
    }

    return d1 == target && d2 == target
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
