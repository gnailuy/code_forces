package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, t int
	if _, err := fmt.Scanf("%d %d\n", &n, &t); err != nil {
		return
	}

	a := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < n; i++ {
		ai, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		a[i] = ai
	}

	max_books := 0

	start := -1
	time := 0
	books := 0
	for i := 0; i < n; i++ {
		if a[i] > t {
			if books > max_books {
				max_books = books
			}
			start = -1
			time = 0
			books = 0
			continue
		}
		if start < 0 && a[i] <= t {
			start = i
			time += a[i]
			books++
		} else if start >= 0 {
			if time+a[i] <= t {
				time += a[i]
				books++
			} else {
				if books > max_books {
					max_books = books
				}
				for j := start; j < i; j++ {
					time -= a[j]
					books--
					if time+a[i] <= t {
						start = j + 1
						time += a[i]
						books++
						break
					}
				}
			}
		}
	}
	if books > max_books {
		max_books = books
	}

	fmt.Println(max_books)
}
