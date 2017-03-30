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

	pos_count := 0
	neg_count := 0
	zero_count := 0

	array := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if a < 0 {
			array[neg_count] = a
			neg_count++
		} else if a > 0 {
			array[n-pos_count-1] = a
			pos_count++
		} else {
			zero_count++
		}
	}

	if pos_count > 0 {
		if neg_count%2 != 0 {
			fmt.Printf("%d ", neg_count)
			for i := 0; i < neg_count; i++ {
				fmt.Printf("%d ", array[i])
			}
			fmt.Println()
		} else {
			fmt.Printf("%d ", neg_count-1)
			for i := 0; i < neg_count-1; i++ {
				fmt.Printf("%d ", array[i])
			}
			fmt.Println()
		}
		fmt.Printf("%d ", pos_count)
		for i := 0; i < pos_count; i++ {
			fmt.Printf("%d ", array[n-i-1])
		}
		fmt.Println()
		if neg_count%2 != 0 {
			fmt.Printf("%d ", zero_count)
			for i := 0; i < zero_count; i++ {
				fmt.Printf("0 ")
			}
			fmt.Println()
		} else {
			fmt.Printf("%d ", zero_count+1)
			for i := 0; i < zero_count; i++ {
				fmt.Printf("0 ")
			}
			fmt.Println(array[neg_count-1])
		}
	} else {
		if neg_count%2 != 0 {
			fmt.Printf("%d ", neg_count-2)
			for i := 0; i < neg_count-2; i++ {
				fmt.Printf("%d ", array[i])
			}
			fmt.Println()
		} else {
			fmt.Printf("%d ", neg_count-3)
			for i := 0; i < neg_count-3; i++ {
				fmt.Printf("%d ", array[i])
			}
			fmt.Println()
		}
		fmt.Printf("2 ")
		for i := 0; i < 2; i++ {
			fmt.Printf("%d ", array[neg_count+i-2])
		}
		fmt.Println()
		if neg_count%2 != 0 {
			fmt.Printf("%d ", zero_count)
			for i := 0; i < zero_count; i++ {
				fmt.Printf("0 ")
			}
			fmt.Println()
		} else {
			fmt.Printf("%d ", zero_count+1)
			for i := 0; i < zero_count; i++ {
				fmt.Printf("0 ")
			}
			fmt.Println(array[neg_count-3])
		}
	}
}
