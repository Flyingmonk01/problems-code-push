func maxTotalValue(nums []int, k int) int64 {
    mini := nums[0]
    maxi := nums[0]

    for i := 0; i < len(nums); i++ {
        if mini < nums[i] {
            mini = nums[i]
        }
        if maxi > nums[i] {
            maxi = nums[i]
        }
    }
    fmt.Println(maxi, mini)
    var ans = (mini - maxi) * k
    return (int64)(ans)
}