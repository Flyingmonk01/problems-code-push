func isSorted(nums []int, n int) bool {
    for i:=1; i < n; i++ {
        if nums[i] < nums[i-1]{
            return false
        }
    }
    return true
}

func minimumPairRemoval(nums []int) int {
    ans, n := 0, len(nums)

    for isSorted(nums, n) == false {
        ans += 1
        min_sum := int(1e18)
        pos := -1
        for i:=1; i < n; i++ {
            sum := nums[i] + nums[i-1]
            if sum < min_sum {
                min_sum = sum
                pos = i
            }
        }
        nums[pos-1] = min_sum
        for i:=pos; i < n-1; i++ {
            nums[i] = nums[i+1]
        }
        n--
    }
    return ans
}