package main

import (
	"fmt"
)

type AppleTree struct {
	location int
	apples   int
}

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	array := make([]*AppleTree, n)
	neg_cnt := 0
	neg_apple := 0
	pos_cnt := 0
	pos_apple := 0

	for i := 0; i < n; i++ {
		var x, a int
		if _, err := fmt.Scanf("%d %d\n", &x, &a); err != nil {
			return
		}
		array[i] = &AppleTree{x, a}
		if x < 0 {
			neg_cnt++
			neg_apple += a
		} else {
			pos_cnt++
			pos_apple += a
		}
	}

	if neg_cnt == pos_cnt {
		fmt.Println(neg_apple + pos_apple)
	} else {
		apples := 0
		bubble_sort(array, n)

		if neg_cnt == 0 {
			fmt.Println(array[0].apples)
			return
		} else if pos_cnt == 0 {
			fmt.Println(array[n-1].apples)
			return
		} else if neg_cnt < pos_cnt {
			apples += neg_apple
			for i := neg_cnt; i < neg_cnt+neg_cnt+1; i++ {
				apples += array[i].apples
			}
		} else if neg_cnt > pos_cnt {
			apples += pos_apple
			for i := neg_cnt - pos_cnt - 1; i < neg_cnt; i++ {
				apples += array[i].apples
			}
		}

		fmt.Println(apples)
	}
}

func swap(a, b int, array []*AppleTree) {
	tmp := array[a]
	array[a] = array[b]
	array[b] = tmp
}

func bubble_sort(array []*AppleTree, length int) {
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if array[i].location < array[j].location {
				swap(i, j, array)
			}
		}
	}
}
