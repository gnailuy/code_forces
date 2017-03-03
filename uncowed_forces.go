package main

import (
	"fmt"
)

func main() {
	x := [5]int{500, 1000, 1500, 2000, 2500}

	var m [5]int
	for i := 0; i < 4; i++ {
		if _, err := fmt.Scanf("%d", &m[i]); err != nil {
			return
		}
	}
	if _, err := fmt.Scanf("%d\n", &m[4]); err != nil {
		return
	}
	var w [5]int
	for i := 0; i < 4; i++ {
		if _, err := fmt.Scanf("%d", &w[i]); err != nil {
			return
		}
	}
	if _, err := fmt.Scanf("%d\n", &w[4]); err != nil {
		return
	}
	var h_s, h_u int
	if _, err := fmt.Scanf("%d %d\n", &h_s, &h_u); err != nil {
		return
	}

	score := 0
	for i := 0; i < 5; i++ {
		s1 := x[i] * 3 / 10
		s2 := x[i] - m[i]*(x[i]/250) - 50*w[i]
		if s1 > s2 {
			score += s1
		} else {
			score += s2
		}
	}
	score += 100 * h_s
	score -= 50 * h_u

	fmt.Println(score)
}
