package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	if garland, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		garland = re.ReplaceAllString(garland, "")

		needs := make([]int, 4)
		color_idx := make([]int, 4) // Index of RBYG

		for i := 0; i < len(garland); i++ {
			k := i % 4
			if 'R' == garland[i] {
				color_idx[0] = k
			} else if 'B' == garland[i] {
				color_idx[1] = k
			} else if 'Y' == garland[i] {
				color_idx[2] = k
			} else if 'G' == garland[i] {
				color_idx[3] = k
			} else if '!' == garland[i] {
				needs[k]++
			} else {
				fmt.Println("Error")
				return
			}
		}

		for i := 0; i < 4; i++ {
			fmt.Printf("%d ", needs[color_idx[i]])
		}
		fmt.Println()
	}
}
