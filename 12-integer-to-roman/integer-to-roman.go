func getRomanRep(num int) string {
    romanMap := map[int]string{
        1:    "I",
        2:    "II",
        3:    "III",
        4:    "IV",
        5:    "V",
        6:    "VI",
        7:    "VII",
        8:    "VIII",
        9:    "IX",

        10:   "X",
        20:   "XX",
        30:   "XXX",
        40:   "XL",
        50:   "L",
        60:   "LX",
        70:   "LXX",
        80:   "LXXX",
        90:   "XC",

        100:  "C",
        200:  "CC",
        300:  "CCC",
        400:  "CD",
        500:  "D",
        600:  "DC",
        700:  "DCC",
        800:  "DCCC",
        900:  "CM",

        1000: "M",
        2000: "MM",
        3000: "MMM",
    }

    return romanMap[num]
}

func intToRoman(num int) string {
    
    idx := 0
    ans := ""
    for num != 0 {
        lastDigit := num % 10
        lastDigit = lastDigit * int(math.Pow(10, float64(idx)))

        currRomanString := getRomanRep(lastDigit)

        ans = currRomanString + ans

        num /= 10
        idx++
    }

    return ans
}