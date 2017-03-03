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

	cut := 0
	occupied := -1
	var x_i, h_i, last_x_i int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < 2*n; i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if i%2 == 0 {
			if i > 0 {
				last_x_i = x_i
			}
			x_i = a
		} else {
			h_i = a
			if occupied < 0 {
				occupied = x_i
				cut++
			} else if occupied >= x_i {
				occupied = last_x_i
				cut--
				if x_i-h_i > occupied {
					occupied = x_i
					cut++
				} else {
					occupied = x_i + h_i
					cut++
				}
			} else if x_i-h_i > occupied {
				occupied = x_i
				cut++
			} else {
				occupied = x_i + h_i
				cut++
			}
		}
	}

	fmt.Println(cut)
}
