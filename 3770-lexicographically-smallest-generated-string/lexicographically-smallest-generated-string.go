
func isSame(word []byte, str string, i int, m int) bool {
    for j := 0; j < m; j++ {
        if word[i] != str[j] {
            return false
        }
        i++
    }
    return true
}

func generateString(str1 string, str2 string) string {
    n := len(str1)
    m := len(str2)

    N := n + m - 1

    canChange := make([]bool, N)
    word := make([]byte, N)
    for i := 0; i < N; i++ {
        word[i] = '$'
    }

    for i := 0; i < n; i++ {
        if str1[i] == 'T' {
            i_ := i
            for j := 0; j < m; j++ {
                if word[i_] != '$' && word[i_] != str2[j] {
                    return ""
                }
                word[i_] = str2[j]
                i_++
            }
        }
    }

    for i := 0; i < N; i++ {
        if word[i] == '$' {
            word[i] = 'a'
            canChange[i] = true
        }
    }

    for i := 0; i < n; i++ {
        if str1[i] == 'F' {
            if isSame(word, str2, i, m) {

                changed := false
                for k := i + m - 1; k >= i; k-- {
                    if canChange[k] == true {
                        word[k] = 'b'
                        changed = true
                        break
                    }
                }
                if changed == false {
                    return ""
                }
            }
        }
    }
    return string(word)
}