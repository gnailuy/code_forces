package main

import (
	"fmt"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	min := make([]int, 3)
	max := make([]int, 3)

	for i := 0; i < 3; i++ {
		if _, err := fmt.Scanf("%d %d\n", &min[i], &max[i]); err != nil {
			return
		}
		if max[i] > n {
			max[i] = n
		}
	}

	var f, s, t int
	var ok bool = false

	fmi := n - max[1] - max[2]
	if fmi < 1 || fmi > n-2 || fmi < min[0] {
		fmi = min[0]
	}
	fmx := n - min[1] - min[2]
	if fmx < 1 || fmx > n-2 || fmx > max[0] {
		fmx = max[0]
	}
	for i := fmx; i >= fmi; i-- {
		smi := n - i - max[2]
		if smi < 1 || smi > n-2 || smi < min[1] {
			smi = min[1]
		}
		smx := n - i - min[2]
		if smx < 1 || smx > n-2 || smx > max[1] {
			smx = max[1]
		}
		for j := smx; j >= smi; j-- {
			k := n - i - j
			if k >= min[2] && k <= max[2] {
				f = i
				s = j
				t = k
				ok = true
				break
			}
		}
		if ok {
			break
		}
	}

	fmt.Printf("%d %d %d\n", f, s, t)
}
