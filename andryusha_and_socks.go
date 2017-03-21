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

	table := make(map[int]bool)
	table_count := 0
	max_table_count := 0

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < 2*n && scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}

		if _, ok := table[a]; ok {
			delete(table, a)
			table_count -= 1
		} else {
			table[a] = true
			table_count += 1
			if table_count > max_table_count {
				max_table_count = table_count
			}
		}
	}

	fmt.Println(max_table_count)
}
