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

	var changed_bits int = 0
	reader := bufio.NewReader(os.Stdin)
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		for i := 0; i < n; i++ {
			if s[i] == '0' {
				changed_bits += 1
				break
			}
			changed_bits += 1
		}
	}

	fmt.Println(changed_bits)
}
