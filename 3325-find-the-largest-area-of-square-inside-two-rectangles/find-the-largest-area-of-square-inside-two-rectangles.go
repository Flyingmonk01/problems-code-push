func largestSquareArea(bL [][]int, tR [][]int) int64 {
    ans := 0
    n := len(bL)
    
    for i:=0; i < n; i++ {
        for j:=i+1; j < n; j++ {
            minX := max(bL[i][0], bL[j][0])
            minY := max(bL[i][1], bL[j][1])
            maxX := min(tR[i][0], tR[j][0])
            maxY := min(tR[i][1], tR[j][1])

            if minX < maxX && minY < maxY {
                curr_size := min(maxX-minX, maxY-minY)
                ans = max(ans, curr_size)
            }
        }
    }

    return int64(ans) * int64(ans)
}