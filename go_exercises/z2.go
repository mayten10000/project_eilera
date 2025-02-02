package main

import "fmt"

func main() {
	limit := 4000000
	sum := 0

	a, b := 1, 2
	for b < limit {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	fmt.Println("Сумма чисел ", sum)
}
