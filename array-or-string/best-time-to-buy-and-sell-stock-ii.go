// can buy at same day, but only hold one stock
// on the v 1, only buy and sell (one time transaction)
// in here we can over buy and sell if it can make profit
// so actually with this problem, the profit must be high than v1

// we use greedy, if next day price is higher, buy and then sell
// add it to profit

func maxProfit(prices []int) int {
	var profit int = 0
	for i := 0; i < len(prices)-1; i++ {
		if potentialProfit := prices[i+1] - prices[i]; potentialProfit > 0 {
			profit += potentialProfit
		}
	}

	return profit
}

// complexity time : O(n), space : O(1)