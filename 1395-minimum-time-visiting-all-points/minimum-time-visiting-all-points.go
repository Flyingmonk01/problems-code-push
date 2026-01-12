func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func minTimeToVisitAllPoints(points [][]int) int {
    
    ans := 0
    for i:=1; i < len(points); i++ {
        point := points[i]
        prev_point := points[i-1]

        x := abs(point[0]-prev_point[0])
        y := abs(point[1]-prev_point[1])

        ans = ans + max(x, y)
    }

    return ans
}