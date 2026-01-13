
func min(a int, b int) int {
    if a < b{
        return a
    }
    return b
}

func max(a int, b int) int {
    if a < b {
        return b
    }
    return a
}

func separateSquares(squares [][]int) float64 {
    var totalArea float64 = 0
    minY := math.MaxInt
    maxY := math.MinInt

    for _, arr := range squares {
        totalArea += (float64)(arr[2]*arr[2])
        minY = min(minY, arr[1])
        maxY = max(maxY, arr[1]+arr[2])
    }

    target := totalArea / 2
    low := float64(minY)
    high := float64(maxY)

    for i := 0; i < 60; i++ {
        mid := (low + high) / 2.0
		var areaBelow float64 = 0

        for _, sq := range squares {
            y := float64(sq[1])
			l := float64(sq[2])

            if mid <= y {
				continue
			} else if mid >= y+l {
				areaBelow += l * l
			} else {
				areaBelow += l * (mid - y)
			}
        }
        if areaBelow < target {
			low = mid
		} else {
			high = mid
		}
    }

    return (low + high) / 2

}