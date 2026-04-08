const MOD int = 1000000007

func xorAfterQueries(nums []int, queries [][]int) int {
    res := 0

    for _, curr_query := range queries {
        start := curr_query[0]
        end := curr_query[1]

        for start <= end && start < len(nums){
            nums[start] = (nums[start] * curr_query[3]) % MOD
            start += curr_query[2]
        }
    }

    for i := 0; i < len(nums); i++ {
        res ^= nums[i]
    }

    return res
}