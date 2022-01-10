// Runtime: 112 ms, faster than 99.11% of Go online submissions for Best Time to Buy and Sell Stock.
// Memory Usage: 8.3 MB, less than 96.66% of Go online submissions for Best Time to Buy and Sell Stock.

func maxProfit(prices []int) int {
    maxProfit := 0
    bestBuy := 1<<31 - 1 // max int in 4 bytes
    
    for _, p := range prices {
        if p < bestBuy {
            bestBuy = p
        } else if p - bestBuy > maxProfit{
            maxProfit = p - bestBuy
        } 
    }
    
    return maxProfit
    
}