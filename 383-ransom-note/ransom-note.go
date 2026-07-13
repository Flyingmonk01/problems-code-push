func canConstruct(ransomNote string, magazine string) bool {
    freq := make([]int, 26)

    for i := 0; i < len(ransomNote); i++ {
        freq[ransomNote[i]-'a']++
    }
    for i := 0; i < len(magazine); i++ {
        if freq[magazine[i]-'a'] > 0 {
            freq[magazine[i]-'a']--
        }
    }
    for i := 0; i < 26; i++ {
        if freq[i] > 0 {
            return false
        }
    }
    return true
}