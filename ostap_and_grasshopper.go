package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var n, k int
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	var ok bool = false

	reader := bufio.NewReader(os.Stdin)
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		var pos_grasshopper int
		var pos_insect int

		for i := 0; i < n; i++ {
			if s[i] == 'G' {
				pos_grasshopper = i
			} else if s[i] == 'T' {
				pos_insect = i
			}
		}
		if pos_grasshopper > pos_insect {
			k = -k
		}
		for i := pos_grasshopper + k; i >= 0 && i < n; i += k {
			if s[i] == 'T' {
				ok = true
				break
			} else if s[i] == '#' {
				ok = false
				break
			}
			if (k > 0 && i > pos_insect) || (k < 0 && i < pos_insect) {
				ok = false
				break
			}
		}
	}

	if ok {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
