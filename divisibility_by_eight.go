package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\r?\n`)
	reader := bufio.NewReader(os.Stdin)
	if in, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		in = re.ReplaceAllString(in, "")
		n := len(in)

		if n == 1 {
			d := int(in[0] - '0')
			if d%8 == 0 {
				fmt.Println("YES")
				fmt.Println(d)
			} else {
				fmt.Println("NO")
			}
		} else {
			digits := make([]int, n)
			dp := make([][]int, n)
			store := make([][]string, n)
			for i := 0; i < n; i++ {
				digits[i] = int(in[i] - '0')
				dp[i] = make([]int, 8)
				store[i] = make([]string, 8)
			}

			dp[0][digits[0]%8] = 1
			store[0][digits[0]%8] = string(in[0])
			for i := 1; i < n; i++ {
				for j := 0; j < 8; j++ {
					dp[i][digits[i]%8] = 1 // If all previous digits are not added
					store[i][digits[i]%8] = string(in[i])
					if dp[i-1][j] == 1 {
						// Adding this digit
						dp[i][(j*10+digits[i])%8] = 1
						store[i][(j*10+digits[i])%8] = store[i-1][j] + string(in[i])
						// Deleting this digit
						dp[i][j] = 1
						store[i][j] = store[i-1][j]
					}
				}
			}

			ok := false
			for i := 1; i < n; i++ {
				if dp[i][0] == 1 && len(store[i][0]) > 0 {
					fmt.Println("YES")
					fmt.Println(store[i][0])
					ok = true
					break
				}
			}
			if !ok {
				fmt.Println("NO")
			}
		}
	}
}
