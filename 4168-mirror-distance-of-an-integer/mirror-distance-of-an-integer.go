func getReverse(n int) int {
    m := 0

    for n != 0 {
        m = m*10 + n % 10
        n = n / 10
    }

    return m
}

func mirrorDistance(n int) int {
    m := getReverse(n)
    if m > n {
        return m-n
    }
    return n-m
}