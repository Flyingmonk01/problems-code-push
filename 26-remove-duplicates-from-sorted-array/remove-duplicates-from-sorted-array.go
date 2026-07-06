func removeDuplicates(nums []int) int {
    i, j := 0, 0

    for i < len(nums) {
        if nums[i] == nums[j] {
            i++
        }else {
            j++
            nums[j] = nums[i]
        }
    }
    return j+1
}