
func maxDotProductRecursion(nums1 []int, nums2 []int, i int, j int, dp [][]int) int {
    if i >= len(nums1) || j >= len(nums2) {
        return 0
    }

    if dp[i][j] != -1 {
        return dp[i][j]
    }

    var inc int = nums1[i]*nums2[j] + maxDotProductRecursion(nums1, nums2, i+1, j+1, dp)

    var exc int = max(maxDotProductRecursion(nums1, nums2, i+1, j, dp), maxDotProductRecursion(nums1, nums2, i, j+1, dp))

    dp[i][j] = max(inc, exc)
    return dp[i][j]
}

func maxDotProduct(nums1 []int, nums2 []int) int {
    min1, min2, max1, max2 := 1000, 1000, -1000, -1000

    dp := make([][]int, len(nums1)+1)
    for i:=0; i <= len(nums1); i++ {
        dp[i] = make([]int, len(nums2)+1)
        for j:=0; j <= len(nums2); j++ {
            dp[i][j] = -1
        }
    }

    for _,i := range nums1 {
        min1 = min(min1, i)
        max1 = max(max1, i)
    }
    for _,i := range nums2 {
        min2 = min(min2, i)
        max2 = max(max2, i)
    }

    if min1 > 0 && max2 < 0 {
        return min1*max2
    }
    if min2 > 0 && max1 < 0 {
        return min2*max1
    }

    return maxDotProductRecursion(nums1, nums2, 0, 0, dp)
}