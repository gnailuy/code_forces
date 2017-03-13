package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	xs := make([]int, n)
	ys := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		scanner.Scan()
		y, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		xs[i] = x
		ys[i] = y
	}

	cnt := 0
	for i := 0; i < n; i++ {
		if is_supercental(xs, ys, i, n) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func is_supercental(xs, ys []int, p, n int) bool {
	array := [4]bool{false, false, false, false} // Upper, Lower, Right, Left
	for i := 0; i < n; i++ {
		if i != p {
			if xs[i] == xs[p] {
				if ys[i] > ys[p] {
					array[0] = true
				} else if ys[i] < ys[p] {
					array[1] = true
				}
			}
			if ys[i] == ys[p] {
				if xs[i] > xs[p] {
					array[2] = true
				} else if xs[i] < xs[p] {
					array[3] = true
				}
			}
		}
	}
	if array[0] && array[1] && array[2] && array[3] {
		return true
	} else {
		return false
	}
}
