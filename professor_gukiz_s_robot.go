package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 int
	if _, err := fmt.Scanf("%d %d\n", &x1, &y1); err != nil {
		return
	}
	if _, err := fmt.Scanf("%d %d\n", &x2, &y2); err != nil {
		return
	}

	diffx := math.Abs(float64(x1 - x2))
	diffy := math.Abs(float64(y1 - y2))

	fmt.Println(int(math.Max(diffx, diffy)))
}
