package main

import (
	"fmt"
	m "math"
)

func main() {
	var r, x, y, x_dash, y_dash int
	if _, err := fmt.Scanf("%d %d %d %d %d\n",
		&r, &x, &y, &x_dash, &y_dash); err != nil {
		return
	}

	var distance = m.Sqrt(m.Abs(float64(x-x_dash))*m.Abs(float64(x-x_dash)) +
		m.Abs(float64(y-y_dash))*m.Abs(float64(y-y_dash)))
	var moves = int(distance / (float64(2 * r)))
	if distance-float64(moves*r*2) > 0 {
		moves++
	}

	fmt.Printf("%d\n", moves)
}
