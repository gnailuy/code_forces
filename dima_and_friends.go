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

	fingers := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		fingers += a
	}

	ways := 0
	for i := 1; i <= 5; i++ {
		if (fingers+i)%(n+1) != 1 {
			ways++
		}
	}

	fmt.Println(ways)
}
