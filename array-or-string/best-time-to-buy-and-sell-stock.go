// Kadane algorithm
// choose one day to buy, and one day to sell (only one transaction)
// iterate through array with save the minBuyPrice
// if now price is lower than curr minBuy, then change it
// if now price is higher, than calculate the profit
// will calculate profit based on minBuyPrice and the current price
// if now price mines minBuyPrice > curr profit, then change it
// return profit

func maxProfit(prices []int) int {
	var profit = 0
	var minPrice = prices[0]

	for i := 1; i < len(prices); i++ {
		// If we find any price which is lower than the current minPrice
		// update the minPrice
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > profit {
			// If diff of current stock with minPrice is greater
			// update the profit
			profit = prices[i] - minPrice
		}
	}

	return profit
}

// complexity time : O(n), space : O(1)