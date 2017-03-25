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

	steps := int64(0)
	current_value := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if a != current_value {
			steps += int64(abs(a - current_value))
			current_value = a
		}
	}

	fmt.Println(steps)
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
