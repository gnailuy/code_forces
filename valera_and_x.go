package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var n int // Odd
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	var is_x bool = true
	var diagonal byte
	var blank byte

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
				if i == 0 && j == 0 {
					diagonal = s[j]
				} else if i == 0 && j == 1 {
					blank = s[j]
					if blank == diagonal {
						is_x = false
					}
				} else {
					if j == i || j == n-i-1 {
						if s[j] != diagonal {
							is_x = false
						}
					} else {
						if s[j] != blank {
							is_x = false
						}
					}
				}
			}
		}
	}

	if is_x {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
