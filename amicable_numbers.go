package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	array := make([]int64, n)
	res := int64(0)

	for i := 1; i <= n; i++ {
		if array[i-1] == 0 {
			array[i-1] = sum_of_divisors(i)
		} else if array[i-1] < 0 {
			continue
		}
		di := array[i-1]
		if di != int64(i) && di <= int64(n) {
			if array[di-1] == 0 {
				array[di-1] = sum_of_divisors(int(di))
			}
			ddi := array[di-1]
			if int64(i) == ddi {
				fmt.Println("Adding:", i, di)
				res += (int64(i) + di)
			}
		}
	}

	fmt.Println(res / 2)
}

func sum_of_divisors(n int) int64 {
	res := int64(1)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			res += int64(i)
			if i != n/i {
				res += int64(n / i)
			}
		}
	}
	return res
}
