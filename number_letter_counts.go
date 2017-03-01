package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\ |\-`)
	total := int64(0)
	for i := 1; i <= 1000; i++ {
		word := build_word(i)
		fmt.Println(word)
		word = re.ReplaceAllString(word, "")
		total += int64(len(word))
	}
	fmt.Println(total)
}

func build_word(i int) string {
	digits := [10]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine"}
	teens := [10]string{"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen"}
	tys := [10]string{"zero",
		"ten",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety"}
	hundred := "hundred"
	thousand := "thousand"

	under_100 := func(i int) string {
		var word string
		if i < 10 {
			word = digits[i]
		} else if i < 20 {
			word = teens[i-10]
		} else if i < 100 {
			if i%10 == 0 {
				word = tys[i/10]
			} else {
				word = tys[i/10] + "-" + digits[i%10]
			}
		} else {
			fmt.Println("under_100() does support", i)
		}
		return word
	}

	under_1000 := func(i int) string {
		var word string
		if i < 100 {
			word = under_100(i)
		} else if i < 1000 {
			word = digits[i/100] + " " + hundred
			j := i % 100
			if j > 0 {
				word += (" and " + under_100(j))
			}
		} else {
			fmt.Println("under_1000() does support", i)
		}
		return word
	}

	under_10000 := func(i int) string {
		var word string
		if i < 1000 {
			word = under_1000(i)
		} else if i < 10000 {
			word = digits[i/1000] + " " + thousand
			j := i % 1000
			if j > 0 {
				word += (", " + under_1000(j))
			}
		} else {
			fmt.Println("under_10000() does support", i)
		}
		return word
	}

	var word string
	if i < 10000 {
		word = under_10000(i)
	} else {
		fmt.Println("Unsupported:", i)
		word = ""
	}

	return word
}
