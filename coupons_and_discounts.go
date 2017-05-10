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

	var have_tdc, failed bool = false, false

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if a > 0 {
			a %= 2
			if a > 0 {
				have_tdc = !have_tdc
			}
		} else if have_tdc {
			failed = true
		}
	}

	if failed || have_tdc {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
