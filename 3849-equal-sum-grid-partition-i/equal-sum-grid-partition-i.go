func canPartitionGrid(grid [][]int) bool {
    m, n := len(grid), len(grid[0])

    total_sum := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            total_sum += grid[i][j]
        }
    }

    if total_sum & 1 == 1 {
        return false
    }

    target := total_sum / 2
    curr_sum := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            curr_sum += grid[i][j]
        }
        if curr_sum == target {
            return true
        }
    }

    curr_sum = 0
    for j := 0; j < n; j++ {
        for i := 0; i < m; i++ {
            curr_sum += grid[i][j]
        }
        if curr_sum == target {
            return true
        }
    }

    return false
}