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

	array := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		array[i] = a
	}

	for i := 0; i < n/2; i++ {
		if i%2 == 0 {
			j := n - i - 1
			array[i], array[j] = array[j], array[i]
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()
}
