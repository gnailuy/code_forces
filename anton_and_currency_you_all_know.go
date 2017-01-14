package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	if str, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		var n int
		for n = 0; str[n] != '\n' && str[n] != '\r'; n++ {
		}

		s := make([]byte, n)

		var last_digit byte = str[n-1] - '0'
		var last_even_idx int = -1
		var swapped_idx int = -1

		for i := 0; i < n-1; i++ {
			s[i] = str[i] - '0'
			if s[i]%2 == 0 {
				last_even_idx = i
				if s[i] < last_digit && swapped_idx < 0 {
					s[n-1] = s[i]
					s[i] = last_digit
					swapped_idx = i
				}
			}
		}
		if swapped_idx < 0 {
			if last_even_idx < 0 {
				fmt.Println("-1")
				return
			} else {
				s[n-1] = s[last_even_idx]
				s[last_even_idx] = last_digit
			}
		}
		for i := 0; i < n; i++ {
			fmt.Printf("%d", s[i])
		}
		fmt.Println()
	}
}
