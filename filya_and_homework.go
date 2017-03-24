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

	v := make([]int, 3)
	vi := 0
	overflow := false

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if overflow {
			continue
		}

		exist := false
		for j := 0; j < vi && j < 3; j++ {
			if v[j] == a {
				exist = true
			}
		}

		if !exist {
			if vi < 3 {
				v[vi] = a
				vi++
			} else {
				overflow = true
			}
		}
	}

	if overflow {
		fmt.Println("NO")
	} else if vi < 3 {
		fmt.Println("YES")
	} else {
		sum := int64(v[0]) + int64(v[1]) + int64(v[2])
		if sum%3 != 0 {
			fmt.Println("NO")
		} else {
			avg := sum / 3
			for i := 0; i < 3; i++ {
				if avg == int64(v[i]) {
					fmt.Println("YES")
					return
				}
			}
			fmt.Println("NO")
		}
	}
}
