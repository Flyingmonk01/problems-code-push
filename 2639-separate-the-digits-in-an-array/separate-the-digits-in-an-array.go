func separateDigits(nums []int) []int {
    var ans []int

    for i:=0; i < len(nums); i++ {
        curr_num := nums[i]
        var curr_ans []int
        for curr_num != 0 {
            curr_ans = append(curr_ans, curr_num%10)
            curr_num = curr_num / 10
        }
        // slices.Reverse(curr_ans)
        for j:=len(curr_ans)-1; j >= 0; j-- {
            ans = append(ans, curr_ans[j])
        }
    }

    return ans
}