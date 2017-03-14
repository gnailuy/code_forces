package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	X := int64(0)
	Y := int64(0)

	var n, bx, m, by int
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	scanner.Scan()
	bx, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		a, e := strconv.Atoi(scanner.Text())
		if e != nil {
			return
		}
		X += int64(math.Pow(float64(bx), float64(n-i-1))) * int64(a)
	}

	scanner.Scan()
	m, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	scanner.Scan()
	by, err = strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	for i := 0; i < m; i++ {
		scanner.Scan()
		a, e := strconv.Atoi(scanner.Text())
		if e != nil {
			return
		}
		Y += int64(math.Pow(float64(by), float64(m-i-1))) * int64(a)
	}

	if X == Y {
		fmt.Println("=")
	} else if X > Y {
		fmt.Println(">")
	} else {
		fmt.Println("<")
	}
}
