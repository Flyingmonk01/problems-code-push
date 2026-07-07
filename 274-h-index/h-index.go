func hIndex(c []int) int {
    n := len(c)
    sort.Ints(c)

    for i, val := range c {
        if n - i <= val {
            return n-i
        }
    }
    return 0
}