func reverseBits(n int) int {
    ans := 0

    for i := 1; i <= 32; i++ {
        ans = ans << 1
        ans |= (n&1)
        n = n >> 1
    }
    return ans
}