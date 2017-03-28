package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	if _, err := fmt.Scanf("%d\n", &n); err != nil {
		return
	}

	c := make([][]int, 2)
	c[0] = make([]int, n)
	c[1] = make([]int, n)
	k := make([]int, 2)
	pos := make([]int, 2)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	steps := 0
	winner := -1
	loop := false

	situations := make(map[string]int)

	for j := 0; j < 2; j++ {
		scanner.Scan()
		var ek error
		k[j], ek = strconv.Atoi(scanner.Text())
		if ek != nil {
			return
		}
		for i := 0; i < k[j] && scanner.Scan(); i++ {
			a, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return
			}
			c[j][i] = a
		}
	}

	situations[situation_to_code(c, k, pos, n)] = 1

	for true {
		f := make([]int, 2)
		f[0] = c[0][pos[0]]
		f[1] = c[1][pos[1]]

		for i := 0; i < 2; i++ {
			c[i][pos[i]] = 0
			pos[i]++
			if pos[i] >= n {
				pos[i] = pos[i] % n
			}
		}

		ic := -1
		if f[0] > f[1] {
			ic = 0
		} else { // not equal guaranteed
			ic = 1
		}

		ip := pos[ic] + k[ic] - 1
		if ip >= n {
			ip = ip % n
		}
		c[ic][ip] = f[(ic+1)%2]
		ip++
		if ip >= n {
			ip = ip % n
		}
		c[ic][ip] = f[ic]

		k[ic]++
		k[(ic+1)%2]--

		steps++

		curr_sit := situation_to_code(c, k, pos, n)
		if situations[curr_sit] > 0 {
			loop = true
			break
		} else {
			situations[curr_sit]++
		}
		if k[(ic+1)%2] == 0 {
			winner = ic + 1
			break
		}
	}

	if loop {
		fmt.Println(-1)
	} else {
		fmt.Println(steps, winner)
	}
}

func situation_to_code(c [][]int, k, p []int, n int) string {
	res := make([]string, n+2)
	res[0], res[1] = strconv.Itoa(k[0]), strconv.Itoa(k[1])
	rp := 2
	for i := 0; i < 2; i++ {
		for j := 0; j < k[i]; j++ {
			pos := (p[0] + j) % n
			res[rp] = strconv.Itoa(c[i][pos])
			rp++
		}
	}
	return strings.Join(res, ",")
}
