package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	const max_digits int = 10
	digits := make([]rune, max_digits)
	length := 0
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < max_digits; i++ {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else if c != '4' && c != '7' {
			length = i
			break
		} else {
			digits[i] = c
		}
	}

	var index int = 1 + 2*((1<<uint(length-1))-1) // 1 + 2 + 4 + ... + 2^(length-1)
	for i := 0; i < length; i++ {
		if digits[i] == '7' {
			index += (1 << uint(length-i-1)) // 2^(length-i-1)
		}
	}
	fmt.Println(index)
}
