package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		var i int = 0
		var removed bool = false
		for s[i+1] != '\r' && s[i+1] != '\n' {
			if !removed {
				if s[i] == '0' {
					removed = true
				} else {
					fmt.Printf("%c", s[i])
				}
			} else {
				fmt.Printf("%c", s[i])
			}
			i++
		}
		if removed {
			fmt.Printf("%c", s[i])
		}
		fmt.Println()
	}
}
