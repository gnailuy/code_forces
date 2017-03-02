package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x0, y0 int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &x0, &y0); err != nil {
		return
	}

	slopes := make([][]int, n)
	for i := 0; i < n; i++ {
		var x, y int
		if _, err := fmt.Scanf("%d %d\n", &x, &y); err != nil {
			return
		}
		// numerator, denominator
		numerator := y - y0
		denominator := x - x0
		// numerator and denominator cannot be both 0
		g := gcd(int(math.Abs(float64(numerator))), int(math.Abs(float64(denominator))))
		numerator = numerator / g
		denominator = denominator / g
		// Make sure that denominator is positive or zero
		if numerator <= 0 && denominator <= 0 || numerator > 0 && denominator < 0 {
			numerator = -numerator
			denominator = -denominator
		}
		slopes[i] = make([]int, 2)
		slopes[i][0] = numerator
		slopes[i][1] = denominator
	}

	shots := 0
	for i := 0; i < n; i++ {
		if slopes[i][1] != -1 { // denominator == 0 means stormtrooper destoried
			shots++
			for j := i + 1; j < n; j++ {
				if slopes[j][0] == slopes[i][0] && slopes[j][1] == slopes[i][1] {
					slopes[j][1] = -1 // Mark as destoried
				}
			}
			slopes[i][1] = -1
		}
	}
	fmt.Println(shots)
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
