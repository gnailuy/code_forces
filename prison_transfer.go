package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, t, c int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &t, &c); err != nil {
		return
	}

	plan := 0
	last_ok := -1

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < n; i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if a > t {
			if last_ok >= 0 && i-last_ok >= c {
				plan += i - last_ok - c + 1
			}
			last_ok = -1
		} else {
			if last_ok < 0 {
				last_ok = i
			}
		}
	}
	if last_ok >= 0 && n-last_ok >= c {
		plan += n - last_ok - c + 1
	}

	fmt.Println(plan)
}
