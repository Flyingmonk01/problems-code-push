func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    cnt := 0
    if coins < costs[0] {
        return 0
    }
    for i := 0; i < len(costs); i++ {
        if coins < costs[i] {
            break
        }

        coins -= costs[i]
        cnt++
    }
    return cnt
}