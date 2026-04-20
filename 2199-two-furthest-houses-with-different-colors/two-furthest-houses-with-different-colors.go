func maxDistance(colors []int) int {
    ans := -1
    for i := 0; i < len(colors); i++ {
        for j := len(colors) - 1; j >= i; j-- {
            if colors[i] != colors[j] {
                curr_ans := j - i
                if ans < curr_ans {
                    ans = curr_ans
                }
            }
        }
    }

    return ans
}