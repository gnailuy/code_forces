package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	if n-m > 1 || m > (n+1)*2 {
		fmt.Println(-1)
	} else {
		buffer := make([]string, n+m)
		p := 0
		exm := m - n

		if exm > 0 {
			if m == (n+1)*2 {
				buffer[p] = "1"
				buffer[p+1] = "1"
				p += 2
				exm -= 2
			} else {
				buffer[p] = "1"
				p += 1
				exm -= 1
			}

			for i := 0; i < exm; i++ {
				buffer[p] = "0"
				buffer[p+1] = "1"
				buffer[p+2] = "1"
				p += 3
			}
			for i := 0; i < n-exm; i++ {
				buffer[p] = "0"
				buffer[p+1] = "1"
				p += 2
			}
		} else { // exm is 0 or -1
			for i := 0; i < n-1; i++ {
				buffer[p] = "0"
				buffer[p+1] = "1"
				p += 2
			}
			if exm == 0 {
				buffer[p] = "0"
				buffer[p+1] = "1"
				p += 2
			} else {
				buffer[p] = "0"
				p += 1
			}
		}

		fmt.Println(strings.Join(buffer, ""))
	}
}
