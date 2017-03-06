package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Road struct {
	dest   int
	length int
}

func main() {
	var n, m, k int
	if _, err := fmt.Scanf("%d %d %d\n", &n, &m, &k); err != nil {
		return
	}

	roads := make([][]Road, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		scanner.Scan()
		u, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		scanner.Scan()
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		scanner.Scan()
		l, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		roads[u-1] = append(roads[u-1], Road{v, l})
		roads[v-1] = append(roads[v-1], Road{u, l})
	}

	storage := make(map[int]int)
	for i := 0; i < k; i++ {
		scanner.Scan()
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		storage[a] = a
	}

	if len(storage) < 0 {
		fmt.Println(-1)
		return
	}

	minimal := -1
	for i := 0; i < n; i++ {
		if _, exist := storage[i+1]; !exist { // Non storage city
			for j := 0; j < len(roads[i]); j++ {
				if _, exist := storage[roads[i][j].dest]; exist { // Storage city
					if minimal < 0 || minimal > roads[i][j].length {
						minimal = roads[i][j].length
					}
				}
			}
		}
	}
	fmt.Println(minimal)
}
