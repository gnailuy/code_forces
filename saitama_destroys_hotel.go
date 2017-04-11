package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, s int
	if _, err := fmt.Scanf("%d %d\n", &n, &s); err != nil {
		return
	}

	l := make([][2]int, n+1)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		f, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		scanner.Scan()
		t, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		l[i][0], l[i][1] = f, t
	}
	l[n][0], l[n][1] = 0, 0

	bubble_sort(l, n+1)

	time, floor := 0, s

	for i := 0; i <= n; i++ {
		time += floor - l[i][0]
		floor = l[i][0]
		if time < l[i][1] {
			time = l[i][1]
		}
	}

	fmt.Println(time)
}

func swap(a, b int, array [][2]int) {
	f, t := array[a][0], array[a][1]
	array[a][0], array[a][1] = array[b][0], array[b][1]
	array[b][0], array[b][1] = f, t
}

func bubble_sort(array [][2]int, length int) {
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if array[i][0] > array[j][0] {
				swap(i, j, array)
			}
		}
	}
}
