package array

func maxProfit(prices []int) int {
    if len(prices) == 0 || len(prices) == 1{
        return 0
    }

    var max = prices[0]
    var min = prices[0]
    var maxIndex = 0
    var minIndex = 0
    for i := 0; i < len(prices); i++ {
        if prices[i] > max {
            max = prices[i]
            maxIndex = i
        }

        if prices[i] < min {
            min = prices[i]
            minIndex = i
        }
    }

    var maxProfitTemp int
    if minIndex <= maxIndex {
        maxProfitTemp = max - min
    } else {

        leftMax := maxProfit(prices[:maxIndex+1])
        rightMax := maxProfit(prices[minIndex:])
        midMax := maxProfit(prices[maxIndex+1 : minIndex])

        maxProfitTemp = leftMax

        if maxProfitTemp < midMax {
            maxProfitTemp = midMax
        }

        if maxProfitTemp < rightMax {
            maxProfitTemp = rightMax
        }
    }

    return maxProfitTemp
}

func maxProfit2(prices []int) int {
    var minPrice = prices[0]
    var profit = 0
    for _, price :=range prices {

        if price < minPrice {
            minPrice = price
        }

        if price - minPrice > profit {
            profit = price-minPrice
        }
    }

    return profit
}

