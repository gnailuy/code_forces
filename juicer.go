package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, b, d int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &b, &d); err != nil {
		return
	}

	total := int64(0)
	clean := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if a <= b {
			total += int64(a)
			if total > int64(d) {
				clean++
				total = 0
			}
		}
	}

	fmt.Println(clean)
}
