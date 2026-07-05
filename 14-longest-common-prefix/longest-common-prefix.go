func longestCommonPrefix(strs []string) string {
    var ans strings.Builder

    for i := 0; i < len(strs[0]); i++ {
        curr_ch := strs[0][i]
        match := true

        for j := 1; j < len(strs); j++ {
            if i >= len(strs[j]) || curr_ch != strs[j][i] {
                match = false
                break
            }
        }

        if match == false {
            break;
        }else {
            ans.WriteByte(curr_ch)
        }
    }
    return ans.String()
}