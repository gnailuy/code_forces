package main

import (
	"fmt"
	"math/big"
)

func main() {
	f := [2]big.Int{
		*new(big.Int).SetInt64(int64(1)),
		*new(big.Int).SetInt64(int64(1))}

	fmt.Println("1 :", &f[0])
	fmt.Println("2 :", &f[1])

	for i := 3; ; i++ {
		fn := new(big.Int)
		fn.Add(&f[0], &f[1])
		fmt.Println(i, ":", fn)
		if len(fn.String()) >= 1000 {
			fmt.Println("Ans:", i)
			break
		}
		f[0] = f[1]
		f[1] = *fn
	}
}
