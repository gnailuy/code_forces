package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var n, m int64
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)

	xs := make([]int, n)
	ys := make([]int, n)
	xs_len := 0
	ys_len := 0

	var remain int64 = n * n

	for i := int64(0); i < m; i++ {
		if _in, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			xy := strings.Split(re.ReplaceAllString(_in, ""), " ")
			if len(xy) != 2 {
				return
			}
			x, err := strconv.Atoi(xy[0])
			if err != nil {
				return
			}
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				return
			}

			on_x := false
			on_y := false

			if xs[x-1] == 1 {
				on_x = true
			}
			if ys[y-1] == 1 {
				on_y = true
			}

			if !on_x {
				remain -= (n - int64(ys_len))
			}
			if !on_y {
				remain -= (n - int64(xs_len))
			}
			if !on_x && !on_y {
				remain += 1
			}

			if !on_x {
				xs[x-1] = 1
				xs_len++
			}
			if !on_y {
				ys[y-1] = 1
				ys_len++
			}

			fmt.Printf("%d ", remain)
		}
	}
	fmt.Printf("\n")
}
