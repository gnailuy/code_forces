package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	min_c := 0
	max_c := 0
	min_w := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if a > max_c {
			max_c = a
		}
		if a < min_c || min_c == 0 {
			min_c = a
		}
	}
	for i := 0; i < m && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		if a < min_w || min_w == 0 {
			min_w = a
		}
	}

	v := -1
	for i := max_c; i < min_w; i++ {
		if 2*min_c <= i {
			v = i
			break
		}
	}

	fmt.Println(v)
}
