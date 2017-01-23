package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	s := string(scanner.Text())

	scanner.Scan()
	k, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	ws := make([]int, 26)
	max_w := 0
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		ws[i] = a
		if ws[i] > max_w {
			max_w = ws[i]
		}
	}

	var value int = 0
	idx := 0
	for ; idx < len(s); idx++ {
		value += (ws[int(s[idx]-'a')] * (idx + 1))
	}
	for i := 0; i < k; i++ {
		idx++
		value += max_w * idx
	}

	fmt.Println(value)
}
