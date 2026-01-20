func minBitwiseArray(nums []int) []int {

    ans := make([]int, len(nums))
    for idx, val := range nums {
        curr_ans := -1
        for i:=1; i <= val; i++ {
            if i | (i+1) == val {
                curr_ans = i
                break
            }
        }
        ans[idx] = curr_ans
    }
    return ans
}