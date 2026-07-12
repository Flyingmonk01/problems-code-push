func lengthOfLongestSubstring(s string) int {
    i, j := 0, 0
    charMap := make(map[byte]int)
    n := len(s)
    if n == 0 {
        return 0
    }
    res := 1
    for i < n && j < n {
        if charMap[s[j]] == 0 {
            charMap[s[j]]++
            j++
        } else {
            res = max(res, j-i)
            for charMap[s[j]] > 0 {
                charMap[s[i]]--
                i++
            }
            charMap[s[j]]++
            j++
        }
        res = max(res, j-i)
    }
    return res
}