func maxNumberOfBalloons(text string) int {
    balloonMap := make(map[rune]int)
    for _, ch := range text {
        balloonMap[ch]++
    }
    return min(
        balloonMap['b'],
        balloonMap['a'],
        balloonMap['l']/2,
        balloonMap['o']/2,
        balloonMap['n'],
    )
}