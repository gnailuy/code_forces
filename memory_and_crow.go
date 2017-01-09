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

	var tail_value int = 0
	for i := n - 1; i >= 0; i-- {
		if i%2 == 0 {
			array[i] -= tail_value
			tail_value += array[i]
		} else {
			array[i] += tail_value
			tail_value -= array[i]
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()
}
