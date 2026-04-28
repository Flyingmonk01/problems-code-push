func minOperations(grid [][]int, x int) int {
    arr := make([]int, len(grid)*len(grid[0]))
    i := 0
    for _, xGrid := range grid {
        for _, gridVal := range xGrid {
            arr[i] = gridVal
            i++
        }
    }
    slices.Sort(arr)

    ans := 0
    median := arr[len(arr)/2]

    for _, val := range arr {
        var difference int
        if val > median {
            difference = val - median
        }else{
            difference = median - val
        }
        
        if difference % x != 0 {
            return -1
        }

        ans += (difference) / x
    }

    return ans
}