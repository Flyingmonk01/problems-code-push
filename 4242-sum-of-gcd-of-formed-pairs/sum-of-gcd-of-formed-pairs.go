
func getGcd(a, b int) int {
    if b == 0 {
        return a
    }
    if a < b {
        return getGcd(a, b % a)
    }
    return getGcd(b, a % b)
}

func gcdSum(nums []int) int64 {
    n := len(nums)
    mxi := make([]int, n)
    mxi[0] = nums[0]

    for i := 1; i < n; i++ {
        mxi[i] = max(nums[i], mxi[i-1])
    }

    prefixGcd := make([]int, n)

    for i := 0; i < n; i++ {
        prefixGcd[i] = getGcd(nums[i], mxi[i])
    }

    sort.Ints(prefixGcd)

    i, j := 0, n-1
    var ans int64 = 0

    for i < j {
        curr_gcd := getGcd(prefixGcd[i], prefixGcd[j])
        ans += int64(curr_gcd)
        i++
        j--
    }
    return ans
}