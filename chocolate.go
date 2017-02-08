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

	combination := int64(1)
	gap := 1
	met_first_nut := false
	met_nut := false

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if a == 0 {
			gap += 1
		} else if a == 1 {
			if !met_first_nut {
				met_first_nut = true
			} else {
				combination *= int64(gap)
			}
			met_nut = true
			gap = 1
		}
	}

	if met_nut {
		fmt.Println(combination)
	} else {
		fmt.Println(0)
	}
}
