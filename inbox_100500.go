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

	list := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; scanner.Scan(); i++ {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return
		}
		list[i] = a
	}

	click := 0

	list_view := true
	pointer := 0
	for pointer < n {
		if list_view {
			for pointer < n {
				if list[pointer] == 1 {
					click++
					list[pointer] = 0
					list_view = false
					break
				} else {
					pointer++
				}
			}
			if list_view && click > 0 {
				click--
			}
		} else {
			for pointer < n-1 {
				if list[pointer+1] == 1 {
					click++
					pointer++
					list[pointer] = 0
				} else {
					if pointer >= n-2 {
						pointer++
						break
					} else if list[pointer+2] == 1 {
						click += 2
						pointer += 2
						list[pointer] = 0
					} else {
						click++
						pointer += 2
						list_view = true
						break
					}
				}
			}
			if pointer >= n-1 {
				if list_view && click > 0 {
					click--
				}
				break
			}
		}
	}

	fmt.Println(click)
}
