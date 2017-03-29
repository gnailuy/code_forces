package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	situations := make(map[string]int)

	target := make([]int, n)
	gears := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		gears[i] = a
		target[i] = i
	}

	target_code := gears_to_code(target, n)
	situations[gears_to_code(gears, n)] = 1

	for true {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				gears[i] = (gears[i] + 1) % n
			} else {
				if gears[i] > 0 {
					gears[i] = gears[i] - 1
				} else {
					gears[i] = n - 1
				}
			}
		}

		sit := gears_to_code(gears, n)
		if target_code == sit {
			fmt.Println("Yes")
			return
		} else if situations[sit] > 0 {
			fmt.Println("No")
			return
		} else {
			situations[sit] = 1
		}
	}
}

func gears_to_code(g []int, n int) string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = strconv.Itoa(g[i])
	}
	return strings.Join(res, "")
}
