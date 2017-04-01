package main

import (
	"fmt"
)

func main() {
	var a [6]int
	if _, err := fmt.Scanf("%d %d %d %d %d %d\n",
		&a[0], &a[1], &a[2], &a[3], &a[4], &a[5]); err != nil {
		return
	}

	fmt.Println((a[0]+a[1]+a[2])*(a[0]+a[1]+a[2]) - a[0]*a[0] - a[2]*a[2] - a[4]*a[4])
}
