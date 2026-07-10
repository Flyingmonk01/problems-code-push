func maxArea(height []int) int {
    ans := 0
    i, j := 0, len(height)-1

    // for i := 0; i < len(height); i++ {
    //     for j := i + 1; j < len(height); j++ {
    //         ans = max(ans, min(height[i], height[j])* (j - i))
    //     }
    // }

    for i < j {
        if height[i] < height[j] {
            ans = max(ans, height[i] * (j-i))
            i++
        }else{
            ans = max(ans, height[j] * (j-i))
            j--
        }
    }
    return ans
}