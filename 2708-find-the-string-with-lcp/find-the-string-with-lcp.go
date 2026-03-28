
func checkLCP(word []byte) [][]int {
    n := len(word)

    lcp := make([][]int, n)
    for i := 0; i < n; i++ {
        lcp[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        if word[i] == word[n-1] {
            lcp[i][n-1] = 1
        }
    }
    for j := 0; j < n; j++ {
        if word[j] == word[n-1] {
            lcp[n-1][j] = 1
        }
    }

    for i := n-2; i >= 0; i-- {
        for j := n-2; j >= 0; j-- {
            if word[i] == word[j] {
                lcp[i][j] = 1 + lcp[i+1][j+1]
            }else{
                lcp[i][j] = 0
            }
        }
    }

    return lcp
}

func equalMatrix(a, b [][]int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if len(a[i]) != len(b[i]) {
            return false
        }
        for j := range a[i] {
            if a[i][j] != b[i][j] {
                return false
            }
        }
    }
    return true
}

func findTheString(lcp [][]int) string {
    n := len(lcp)

    word := []byte(strings.Repeat("$", n))

    for i := 0; i < n; i++ {

        for j := 0; j < n; j++ {
            if lcp[j][i] != 0 {
                word[i] = word[j]
                break
            }
        }

        if word[i] == '$' {

            forbidden := make([]bool, 26)

            for j := 0; j < i; j++ {
                if lcp[j][i] == 0 && word[j] >= 'a' && word[j] <= 'z'  {
                    forbidden[word[j]-'a'] = true
                }
            }

            for idx := 0; idx < 26; idx++ {
                if forbidden[idx] == false {
                    word[i] = byte(idx + 'a')
                    break
                }
            }

            if word[i] == '$' {
                return ""
            }
        }
    }

    if equalMatrix(checkLCP(word), lcp) {
        return string(word)
    }
    return ""
}