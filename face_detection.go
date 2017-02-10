package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	var n, m int
	if _, err := fmt.Scanf("%d %d\n", &n, &m); err != nil {
		return
	}

	image := make([]string, n)

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
			image[i] = re.ReplaceAllString(s, "")
		}
	}

	num_faces := 0
	for i := 0; i < n-1; i++ {
		for j := 0; j < m-1; j++ {
			face := [4]bool{false, false, false, false}
			for k := i; k <= i+1; k++ {
				for l := j; l <= j+1; l++ {
					if 'f' == image[k][l] {
						face[0] = true
					} else if 'a' == image[k][l] {
						face[1] = true
					} else if 'c' == image[k][l] {
						face[2] = true
					} else if 'e' == image[k][l] {
						face[3] = true
					}
				}
			}
			if face[0] && face[1] && face[2] && face[3] {
				num_faces += 1
			}
		}
	}

	fmt.Println(num_faces)
}
