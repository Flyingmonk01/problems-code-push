func rotateString(s string, goal string) bool {
    n, m := len(s), len(goal)

    if n != m {
        return false
    }
    new_s := s + s
    for i := 0; i < len(new_s) - m; i++ {
        curr_S := new_s[i:i+m]
        if curr_S == goal {
            return true
        }
    }
    return false
}