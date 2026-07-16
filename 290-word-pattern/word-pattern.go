func wordPattern(pattern string, s string) bool {
    wordMap := make(map[string]byte)
    charMap := make(map[byte]string)

    i, j := 0, 0
    n, m := len(s), len(pattern)

    for j < n && i < m {

        if j >= n {
            return false
        }
        start := j
        for j < n && s[j] != ' ' {
            j++
        }

        curr_word := s[start:j]

        if w, ok := charMap[pattern[i]]; ok {
            if w != curr_word {
                return false
            }
        } else {
            charMap[pattern[i]] = curr_word
        }

        if ch, ok := wordMap[curr_word]; ok {
            if ch != pattern[i] {
                return false
            }
        } else {
            wordMap[curr_word] = pattern[i]
        }

        i++
        if j < n {
            j++
        }
    } 

    return i == m && j >= n
}