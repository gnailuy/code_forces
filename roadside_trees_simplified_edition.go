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

	h := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		h[i] = a
	}

	time := h[0] + 1
	height := h[0]
	for i := 1; i < n; i++ {
		if height < h[i] {
			time += 1               // Jump
			time += (h[i] - height) // Climb Up
			height = h[i]           // At the Top
			time += 1               // Eat
		} else if height >= h[i] {
			time += (height - h[i]) // Climb Down
			height = h[i]           // At the Top
			time += 2               // Jump, Eat
		}
	}

	fmt.Println(time)
}
