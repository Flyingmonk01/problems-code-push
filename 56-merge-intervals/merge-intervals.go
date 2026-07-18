func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    ans := [][]int {intervals[0]}

    for _, curr := range intervals[1:] {
        last := ans[len(ans) - 1]

        if last[1] >= curr[0] {
            last[1] = max(last[1], curr[1])
        }else{
            ans = append(ans, curr)
        }
    }
    return ans
}