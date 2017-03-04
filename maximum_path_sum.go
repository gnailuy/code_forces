package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var triangle []int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		triangle = append(triangle, a)
	}

	// Full tree guaranteed
	height := int((math.Sqrt(float64(8*len(triangle)+1)) - 1) / 2)

	for i := height - 1; i >= 1; i-- {
		s := start(i)
		for j := 0; j < i; j++ {
			choosen := max(triangle[s+j+i], triangle[s+j+i+1])
			triangle[s+j] += choosen
			fmt.Printf("%02d ", triangle[s+j])
		}
		fmt.Println()
	}
}

func start(n int) int {
	return (n - 1) * n / 2
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
