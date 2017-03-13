package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	zero := regexp.MustCompile(`0`)

	with_zero := [2]int{0, 0}
	no_zero := [2]int{0, 0}

	for i := 0; i < 2; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else {
			s = re.ReplaceAllString(s, "")
			n, err := strconv.Atoi(s)
			if err != nil {
				return
			}
			with_zero[i] = n

			s = zero.ReplaceAllString(s, "")
			z, err := strconv.Atoi(s)
			if err != nil {
				return
			}
			no_zero[i] = z
		}
	}

	sum := with_zero[0] + with_zero[1]
	sum_string := strconv.Itoa(sum)
	sum_no_zero := no_zero[0] + no_zero[1]

	sz := zero.ReplaceAllString(sum_string, "")
	sum_remove_zero, err := strconv.Atoi(sz)
	if err != nil {
		return
	}

	if sum_no_zero == sum_remove_zero {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
