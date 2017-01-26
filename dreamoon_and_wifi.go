package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	var target int = 0
	var start int = 0
	var question_marks int = 0

	reader := bufio.NewReader(os.Stdin)
	// Command
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		for i := 0; i < len(s); i++ {
			if s[i] == '+' {
				target++
			} else if s[i] == '-' {
				target--
			}
		}
	}

	// Received
	if s, err := reader.ReadString('\n'); err != nil {
		if err != io.EOF {
			return
		}
	} else {
		for i := 0; i < len(s); i++ {
			if s[i] == '+' {
				start++
			} else if s[i] == '-' {
				start--
			} else if s[i] == '?' {
				question_marks++
			}
		}
	}

	var distance int = int(math.Abs(float64(target - start)))
	var prob float64
	if distance > question_marks {
		prob = 0.0
	} else if distance == question_marks {
		prob = math.Pow(0.5, float64(question_marks))
	} else { // distance < question_marks
		if (question_marks-distance)%2 != 0 {
			prob = 0.0
		} else {
			towards := distance + (question_marks-distance)/2
			prob = float64(factorial(question_marks)) /
				float64(factorial(towards)*factorial(question_marks-towards)*
					(1<<uint(question_marks)))
		}
	}

	fmt.Printf("%0.12f\n", prob)
}

func factorial(n int) int64 {
	if n <= 1 {
		return int64(n)
	} else {
		return int64(n) * factorial(n-1)
	}
}
