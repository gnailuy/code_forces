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

	s1 := -1
	e1p := 0
	s2 := -1
	e2 := -1
	has_s3 := false

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if has_s3 {
			continue
		}

		if i == 0 {
			s1 = a
		} else {
			if e2 > a {
				if s2 < 0 {
					e1p = i
					s2 = a
				} else {
					has_s3 = true
				}
			}
		}
		e2 = a
	}

	if s2 < 0 {
		fmt.Println(0)
	} else if has_s3 || e2 > s1 {
		fmt.Println(-1)
	} else {
		fmt.Println(n - e1p)
	}
}
