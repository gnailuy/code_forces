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

	var infos int = 0
	var turns int = 0
	for infos < n {
		for i := 0; i < n; i++ {
			if array[i] >= 0 && array[i] <= infos {
				infos += 1
				array[i] = -1
			}
		}
		if infos < n {
			turns += 1
		} else {
			break
		}
		for i := n - 1; i >= 0; i-- {
			if array[i] >= 0 && array[i] <= infos {
				infos += 1
				array[i] = -1
			}
		}
		if infos < n {
			turns += 1
		} else {
			break
		}
	}

	fmt.Println(turns)
}
