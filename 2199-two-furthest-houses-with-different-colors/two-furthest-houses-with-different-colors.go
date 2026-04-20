func maxDistance(colors []int) int {
    ans := -1
    n := len(colors)

    for j := n-1; j >= 1; j-- {
        if colors[0] != colors[j] {
            if ans < j {
                ans = j
            }
        }
    }

    for j := 0; j < n-1; j++ {
        if colors[n-1] != colors[j] {
            if ans < n-j-1 {
                ans = n-j-1
            }
        }
    }

    return ans
}