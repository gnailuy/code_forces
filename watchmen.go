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
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	pair_count := int64(0)
	s_map := make(map[string]int)
	x_map := make(map[int]map[int]int)
	y_map := make(map[int]map[int]int)

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	for i := 0; i < n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			s = re.ReplaceAllString(s, "")

			xy := strings.Split(s, " ")
			a, err := strconv.Atoi(xy[0])
			if err != nil {
				return
			}
			b, err := strconv.Atoi(xy[1])
			if err != nil {
				return
			}

			if s_map[s] > 0 {
				pair_count += int64(s_map[s])
			}
			s_map[s] += 1
			if x_map[a] == nil {
				x_map[a] = make(map[int]int)
			}
			x_map[a][b] += 1
			if y_map[b] == nil {
				y_map[b] = make(map[int]int)
			}
			y_map[b][a] += 1
		}
	}

	bf := make([]int, n)
	bl := 0

	for _, xm := range x_map {
		if len(xm) > 1 {
			bl = 0
			for _, c := range xm {
				if bl == 0 {
					bf[0] = c
				} else {
					pair_count += int64(bf[bl-1]) * int64(c)
					bf[bl] = bf[bl-1] + c
				}
				bl++
			}
		}
	}

	for _, ym := range y_map {
		if len(ym) > 1 {
			bl = 0
			for _, c := range ym {
				if bl == 0 {
					bf[0] = c
				} else {
					pair_count += int64(bf[bl-1]) * int64(c)
					bf[bl] = bf[bl-1] + c
				}
				bl++
			}
		}
	}

	fmt.Println(pair_count)
}
