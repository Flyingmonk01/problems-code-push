func minimumCost(cost []int) int {
    sort.Ints(cost)
    ans := 0
    n := len(cost)
    rem := n % 3
    for i := n-1; i >= 0; i-- {
        if i%3 == rem {
            continue
        } 
        ans += cost[i]
    }
    return ans
}