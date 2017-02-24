package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var n int // Even
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(s, "")

		lucky := true
		sum1 := 0
		sum2 := 0
		for i := 0; i < n; i++ {
			digit := int(s[i] - '0')
			if digit == 4 || digit == 7 {
				if i < n/2 {
					sum1 += digit
				} else {
					sum2 += digit
				}
			} else {
				lucky = false
				break
			}
		}
		if lucky && sum1 == sum2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
