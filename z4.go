package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	sign := "positive"
	if x >= 0 {
		sign = "positive"
	} else {
		sign = "negative"
	}
	x = int(math.Abs(float64(x)))
	var reversedDigit int

	for x > 0 {
		lastDigit := x % 10
		reversedDigit = reversedDigit*10 + lastDigit

		x = x / 10
	}

	if sign == "negative" {
		reversedDigit = reversedDigit * -1
	}
	return reversedDigit
}

func main() {
	largestNumber := 0
	slice := []int{}
	for i := 99; i <= 999; i++ {
		for j := 99; j <= 999; j++ {
			s := (i * j)
			c := reverse(s)
			if s == c {
				slice = append(slice, s)
			}
		}
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] > largestNumber {
			largestNumber = slice[i]
		}
	}
	fmt.Printf("Наибольший палиндром: %d\n", largestNumber)
}
