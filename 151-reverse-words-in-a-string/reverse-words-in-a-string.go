func reverse(s string) string {
    b := []byte(s)
    for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
        b[i], b[j] = b[j], b[i]
    }
    return string(b)
}

func reverseWords(s string) string {
    i := len(s)-1
    ans := ""

    for i >= 0 {
        var word []byte
        for i >= 0 && s[i] == ' ' {
            i--
        }
        for i >= 0 && s[i] != ' '  {
            word = append(word, s[i])
            i--
        }
        currStr := string(word)
        currStr = reverse(currStr)
        ans += currStr
        ans += " "
    }

    j := len(ans) - 1
    for j >= 0 && ans[j] == ' ' {
        j--
    }

    return ans[0:j+1]
}