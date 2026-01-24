func minPairSum(nums []int) int {
    sort.Ints(nums)
    n := len(nums) - 1
    maxSum := 0
    for i := 0; i <= n/2; i++ {
        sum := nums[i] + nums[n-i]
        if sum > maxSum {
            maxSum = sum
        }
    }
    return maxSum
}