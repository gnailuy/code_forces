package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

const l int = 26

func main() {
	var n, k int64
	if _, err := fmt.Scanf("%d %d\n", &n, &k); err != nil {
		return
	}

	cnt := make([]int64, l)

	reader := bufio.NewReader(os.Stdin)
	re := regexp.MustCompile(`\r?\n`)
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		s = re.ReplaceAllString(s, "")
		for _, c := range s {
			cnt[int(c-'A')] += 1
		}

		bubble_sort(cnt, l)

		res := int64(0)
		for i := 0; i < l; i++ {
			if k >= cnt[i] {
				res += cnt[i] * cnt[i]
				k -= cnt[i]
			} else {
				res += k * k
				break
			}
		}
		fmt.Println(res)
	}
}

func swap(a, b int, array []int64) {
	array[a], array[b] = array[b], array[a]
}

func bubble_sort(array []int64, length int) {
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if array[i] > array[j] {
				swap(i, j, array)
			}
		}
	}
}
