func checkStrings(s1 string, s2 string) bool {
    n := len(s1)
    even := make([]int, 26)
    odd := make([]int, 26)

    for i := 0; i < n; i++ {
        if i&1 == 1 {
            odd[s1[i] - 'a']++
            odd[s2[i] - 'a']--
        }else{
            even[s1[i] - 'a']++
            even[s2[i] - 'a']--
        }
    }

    for i := 0; i < 26; i++ {
        if even[i] != 0 || odd[i] != 0 {
            return false
        }
    }

    return true
}