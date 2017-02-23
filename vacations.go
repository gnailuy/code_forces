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

	array := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		array[i] = a
	}

	d := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		d[i] = make([]int, 3)
	}

	d[0][0] = 0
	d[0][1] = 0
	d[0][2] = 0

	for i := 1; i <= n; i++ {
		d[i][0] = max(d[i-1])
		if array[i-1] == 1 || array[i-1] == 3 {
			d[i][1] = max([]int{d[i-1][0] + 1, d[i-1][2] + 1})
		}
		if array[i-1] == 2 || array[i-1] == 3 {
			d[i][2] = max([]int{d[i-1][0] + 1, d[i-1][1] + 1})
		}
	}

	fmt.Println(n - max(d[n]))
}

func max(nums []int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}
