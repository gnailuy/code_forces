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

	var fives int = 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		} else if a == 5 {
			fives++
		}
	}
	zeros := n - fives

	if zeros <= 0 {
		fmt.Println(-1)
	} else if fives < 9 {
		fmt.Println(0)
	} else {
		num_fives := fives / 9
		for i := 0; i < num_fives*9; i++ {
			fmt.Printf("5")
		}
		for i := 0; i < zeros; i++ {
			fmt.Printf("0")
		}
		fmt.Println()
	}
}
