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

	drifts := make([][2]int, n)

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

		drifts[i] = [2]int{x - 1, y - 1}
	}

	sets := make(map[int]map[[2]int]bool)
	sets[1] = make(map[[2]int]bool)
	sets[1][drifts[0]] = true
	drifts[0][0] = -1
	drifts[0][1] = -1

	for i := 1; i <= n; i++ {
		for true { // Add new drifts into sets[i]
			before := len(sets[i])
			for j := 0; j < n; j++ {
				if drifts[j][0] >= 0 {
					x := drifts[j][0]
					y := drifts[j][1]
					ok := false
					for p, _ := range sets[i] {
						if p[0] == x || p[1] == y {
							ok = true
							break
						}
					}
					if ok {
						sets[i][[2]int{x, y}] = true
						drifts[j][0] = -1
						drifts[j][1] = -1
					}
				}
			}
			if len(sets[i]) == before {
				break
			}
		}

		more := false
		for j := 0; j < n; j++ {
			if drifts[j][0] >= 0 {
				sets[i+1] = make(map[[2]int]bool)
				sets[i+1][drifts[j]] = true
				more = true
				break
			}
		}
		if !more {
			break
		}
	}

	fmt.Println(len(sets) - 1)
}
