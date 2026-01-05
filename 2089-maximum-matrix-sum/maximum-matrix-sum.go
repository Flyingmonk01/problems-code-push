func maxMatrixSum(matrix [][]int) int64 {
    numOfNegative, totalAbsSum, smallestNum := int64(0),int64(0),int64(1_00_000)
    for _, row := range matrix {
        for _, col := range row {
            absVal := int64(math.Abs(float64(col)))
            totalAbsSum += absVal
            if col < 0 {
                numOfNegative++
            }
            if absVal < smallestNum {
                smallestNum = absVal
            }
        }
    }

    if numOfNegative%2 == 1 {
        return totalAbsSum - 2*smallestNum
    }
    return totalAbsSum
}