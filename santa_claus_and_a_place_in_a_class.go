package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &m, &k); err != nil {
		return
	}

	lane := (k-1)/(2*m) + 1
	lane_start := (lane-1)*2*m + 1
	desk := (k-lane_start)/2 + 1
	var place string
	if (k-lane_start)%2 == 0 {
		place = "L"
	} else {
		place = "R"
	}

	fmt.Println(lane, desk, place)
}
