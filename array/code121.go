package array

func maxProfit122(prices []int) int {

    totalProfit := 0
    profit := 0
    lowPrice := prices[0]

    for _, price := range prices {

        if price-lowPrice > profit {
            profit = price - lowPrice
        } else {
            totalProfit += profit
            lowPrice = price
            profit = 0
        }
    }

    totalProfit += profit
    return totalProfit
}

func maxProfit122_2(prices []int) int {

    profit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i]-prices[i-1] > 0 {
            profit += prices[i] - prices[i-1]
        }
    }

    return profit
}
