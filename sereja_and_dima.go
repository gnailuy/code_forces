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

	var sereja_dima [2]int
	var i int = 0 // Sereja first
	var head int = 0
	var tail int = n - 1
	for head <= tail {
		if array[head] > array[tail] {
			sereja_dima[i] += array[head]
			head++
		} else {
			sereja_dima[i] += array[tail]
			tail--
		}
		i = (i + 1) % 2
	}

	fmt.Printf("%d %d\n", sereja_dima[0], sereja_dima[1])
}
