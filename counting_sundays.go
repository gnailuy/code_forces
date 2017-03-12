package main

import (
	"fmt"
)

func main() {
	md := [12]int{
		31, // Jan
		28, // Feb
		31, // Mar
		30, // Api
		31, // May
		30, // Jun
		31, // Jul
		31, // Aug
		30, // Sep
		31, // Oct
		30, // Nov
		31} // Dec

	start := 1900
	end := 2001

	leap := 0
	days_from_19000101 := 0
	sunday_at_0101 := 0

	for year := start; year < end; year++ {
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			leap = 1
		} else {
			leap = 0
		}
		for month := 0; month < 12; month++ {
			if year > 1900 && (days_from_19000101+1)%7 == 0 {
				sunday_at_0101++
			}
			days_from_19000101 += md[month]
			if month == 1 { // Feb
				days_from_19000101 += leap
			}
		}
	}

	fmt.Println(sunday_at_0101)
}
