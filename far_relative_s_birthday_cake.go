package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	cake := make([][]byte, n)
	for i := 0; i < n; i++ {
		cake[i] = make([]byte, n)
	}

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			for j := 0; j < n; j++ {
				if s[j] == 'C' {
					cake[i][j] = 1
				}
			}
		}
	}

	var pairs int = 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			cnt += int(cake[i][j])
		}
		if cnt > 1 {
			pairs += (cnt * (cnt - 1) / 2)
		}
	}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			cnt += int(cake[j][i])
		}
		if cnt > 1 {
			pairs += (cnt * (cnt - 1) / 2)
		}
	}

	fmt.Println(pairs)
}
