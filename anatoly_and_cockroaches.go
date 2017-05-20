package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var n int
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

		moves := [2]int{}
		l := [2][]int{make([]int, n), make([]int, n)}

		for i, c := range s {
			if c == 'r' {
				l[0][i] = 1
				l[1][i] = 1
			} else if c == 'b' {
				l[0][i] = 0
				l[1][i] = 0
			}
		}

		disorder_start := [2][2]int{}
		for i := 0; i < n; i++ {
			for j := 0; j < 2; j++ {
				if l[j][i] != (i+j)%2 {
					moves[j] += 1
					swap := false
					if disorder_start[j][(i+1)%2] < n {
						for k := max(disorder_start[j][(i+1)%2], i+1); k < n; k += 2 { // Swap with k
							if l[j][k] != (k+j)%2 {
								l[j][i], l[j][k] = l[j][k], l[j][i]
								swap = true
								for f := k + 2; f < n; f += 2 {
									if l[j][f] != (f+j)%2 {
										disorder_start[j][(i+1)%2] = f
										break
									}
								}
								break
							}
						}
					}
					if !swap { // Change color
						disorder_start[j][(i+1)%2] = n
						l[j][i] = (i + j) % 2
					}
				}
			}
		}

		if moves[0] < moves[1] {
			fmt.Println(moves[0])
		} else {
			fmt.Println(moves[1])
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
