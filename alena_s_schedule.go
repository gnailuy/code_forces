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

	var pairs int = 0
	var breaks int = 0
	var at_home bool = true

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if a == 1 {
			pairs = pairs + breaks + 1
			breaks = 0
			at_home = false
		} else if !at_home {
			if breaks == 1 {
				at_home = true
				breaks = 0
			} else {
				breaks = 1
			}
		}
	}

	fmt.Println(pairs)
}
