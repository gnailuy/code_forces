package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	const max_distance int = 20000

	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	var is_at_np bool = true
	var is_at_sp bool = false
	var distance_from_np int = 0
	var ok bool = true

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		if s, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			} else {
				return
			}
		} else if ok {
			var direction byte
			amount := 0
			for j := 0; j < len(s) && s[j] != '\r' && s[j] != '\n'; j++ {
				if s[j] == ' ' {
					direction = s[j+1]
					break
				} else {
					amount *= 10
					amount += int(s[j] - '0')
				}
			}
			if is_at_np && direction != 'S' {
				ok = false
				break
			} else if is_at_sp && direction != 'N' {
				ok = false
				break
			}
			if direction == 'S' {
				distance_from_np += amount
				is_at_np = false
				if distance_from_np == max_distance {
					is_at_sp = true
				} else if distance_from_np > max_distance {
					ok = false
					break
				}
			} else if direction == 'N' {
				distance_from_np -= amount
				is_at_sp = false
				if distance_from_np == 0 {
					is_at_np = true
				} else if distance_from_np < 0 {
					ok = false
					break
				}
			}
		}
	}

	if ok && distance_from_np == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
