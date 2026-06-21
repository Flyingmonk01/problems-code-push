func maxIceCream(costs []int, coins int) int {
    freq := make([]int, 100001)
    for _, cost := range costs {
        freq[cost]++
    }

    ans := 0

    for price := 1; price <= 100000; price++ {
        if freq[price] == 0 {
            continue
        }

        canBuy := min(freq[price], coins/price)
        ans = ans + canBuy 
        coins = coins - canBuy * price

        if coins < price {
            break
        }
    }

    return ans
}