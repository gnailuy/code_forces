package main

import (
	"fmt"
)

func main() {
	var max_perimeter int = 1000

	var max_solution int = 0
	var perimeter_at_max_solution int = 0

	for perimeter := 3; perimeter <= max_perimeter; perimeter++ {
		solution := 0
		for h := 1; h < perimeter; h++ {
			for a := 1; a < h; a++ {
				b := perimeter - h - a
				if b <= 0 || b > h {
					continue
				}
				if is_right_angle_triangle(a, b, h) {
					solution += 1
					fmt.Println(perimeter, h, a, b)
				}
			}
		}
		if solution > max_solution {
			max_solution = solution
			perimeter_at_max_solution = perimeter
		}
	}

	fmt.Println(perimeter_at_max_solution)
}

func is_right_angle_triangle(a, b, hypotenuse int) bool {
	if a*a+b*b == hypotenuse*hypotenuse {
		return true
	}
	return false
}
