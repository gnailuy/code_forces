package main

import (
	"fmt"
)

func main() {
	var cc rune
	var c, d int
	if _, err := fmt.Scanf("%c%d\n", &cc, &d); err != nil {
		return
	}
	c = int(cc-'a') + 1

	var moves int = 8
	if c == 1 || c == 8 {
		moves -= 3
	}
	if d == 1 || d == 8 {
		moves -= 3
	}
	if (c == 1 || c == 8) && (d == 1 || d == 8) {
		moves += 1
	}

	fmt.Println(moves)
}
