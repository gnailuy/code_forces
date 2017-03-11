package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	current := 0
	capacity := 0
	record := make(map[int]bool)

	for i := 0; i < n; i++ {
		if _in, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			log := strings.Split(re.ReplaceAllString(_in, ""), " ")
			if len(log) != 2 {
				return
			}

			action := log[0]
			id, err := strconv.Atoi(log[1])
			if err != nil {
				return
			}

			if action == "+" {
				current++
				record[id] = true
				if abs(current) > capacity {
					capacity = abs(current)
				}
			} else if action == "-" {
				if _, ok := record[id]; ok {
					current--
					record[id] = false
				} else {
					capacity++
				}
			}

		}
	}

	fmt.Println(capacity)
}

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}
