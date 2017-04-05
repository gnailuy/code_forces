package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, z int
	if _, err := fmt.Scanf("%d %d %d\n", &x, &y, &z); err != nil {
		return
	}

	a := math.Sqrt(float64(x * z / y))
	b := math.Sqrt(float64(x * y / z))
	c := math.Sqrt(float64(y * z / x))

	fmt.Println(4 * int32(a+b+c))
}
