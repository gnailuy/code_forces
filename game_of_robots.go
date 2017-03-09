package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, k int
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	id := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan() && i < n; i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error!")
			return
		}
		id[i] = a
	}

	for i := 0; i <= n; i++ {
		if int64(i)*int64(i+1)/2 >= int64(k) {
			fmt.Println(id[int(int64(k)-int64(i)*int64(i-1)/2-1)])
			break
		}
	}
}
