func judgeCircle(moves string) bool {
    x, y := 0, 0
    for i := 0; i < len(moves); i++ {
        if moves[i] == 'U' {
            y++
        }else if moves[i] == 'D' {
            y--
        }else if moves[i] == 'R' {
            x++
        }else{
            x--
        }
    }

    return x == 0 && y == 0
}