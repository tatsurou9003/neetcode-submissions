func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxPrice := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price - minPrice > maxPrice {
			maxPrice = price - minPrice
		}
	}
	return maxPrice
}
