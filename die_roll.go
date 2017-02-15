package main

import (
	"fmt"
)

func main() {
	var y, w int
	if _, err := fmt.Scanf("%d %d\n", &y, &w); err != nil {
		return
	}

	win := 6 - max(y, w) + 1
	if win == 0 { // Impossible here
		fmt.Println("0/1")
	} else if win == 6 {
		fmt.Println("1/1")
	} else {
		g := gcd(win, 6)
		fmt.Printf("%d/%d\n", win/g, 6/g)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	if y == 0 {
		return x
	} else {
		return gcd(y, x%y)
	}
}
