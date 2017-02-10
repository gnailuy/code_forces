package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, a int
	if _, err := fmt.Scanf("%d %d\n", &n, &a); err != nil {
		return
	}

	crimes := make([]bool, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < n; i++ {
		t, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if t == 1 {
			crimes[i] = true
		} else {
			crimes[i] = false
		}
	}

	catched := 0
	for d := 0; d <= max(a-1, n-a); d++ {
		if d == 0 || d > min(a-1, n-a) {
			if a-d > 0 {
				if crimes[a-d-1] {
					catched++
				}
			} else if a+d <= n {
				if crimes[a+d-1] {
					catched++
				}
			}
		} else {
			if crimes[a-d-1] && crimes[a+d-1] {
				catched += 2
			}
		}
	}

	fmt.Println(catched)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
