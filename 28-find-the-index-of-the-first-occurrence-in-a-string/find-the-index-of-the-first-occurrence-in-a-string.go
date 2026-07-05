func strStr(haystack string, needle string) int {
    hlen := len(haystack)
    nlen := len(needle)

    if nlen > hlen {
        return -1
    }

    if haystack == needle {
        return 0
    }

    for i := 0; i+nlen <= hlen; i++ {
        curr_str := haystack[i:i+nlen]
        if curr_str == needle {
            return i
        }
    }
    return -1
}