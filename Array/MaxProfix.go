package main

import (
	"fmt"
)


func maxProfix(prices []int , fee int) int {
	
	n := len(prices)
	holdStock := -1 * prices[0]
	saleStock := 0
	for i := 1; i < n; i++ {
		previousHoldStock := holdStock
		if holdStock < (saleStock - prices[i]) {
			holdStock = saleStock - prices[i]
		}

		if saleStock < (previousHoldStock + prices[i] - fee) {
			saleStock = previousHoldStock + prices[i] - fee
		}
	}

	return saleStock
}

func main() { 
	fmt.Println(maxProfix([]int {1, 3, 2, 8, 4, 9}, 2))
}

