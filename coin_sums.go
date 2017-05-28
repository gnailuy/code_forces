package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	if n < 0 {
		fmt.Println(0)
	} else if n <= 1 {
		fmt.Println(1)
	} else {
		var coins = []int{1, 2, 5, 10, 20, 50, 100, 200}
		fmt.Println(find_ways(n, coins))
	}
}

func find_ways(n int, coins []int) int {
	l := len(coins)

	if n == 0 || l <= 1 {
		return 1
	}

	largest := -1
	var sub_coins []int
	for i := 0; i < l-1; i++ {
		if coins[i+1] <= n {
			sub_coins = append(sub_coins, coins[i])
		} else {
			largest = coins[i]
			break
		}
	}
	if largest < 0 {
		largest = coins[l-1]
	}
	max := n / largest

	ways := 0
	for i := 0; i <= max; i++ {
		ways += find_ways(n-largest*i, sub_coins)
	}

	return ways
}
