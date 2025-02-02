package main

import (
	"fmt"
	"math"
)

func main() {
	number := int64(600851475143)
	limit := int64(math.Sqrt(float64(number)))
	for i := int64(2); i < limit; i++ {
		if number%i == 0 {
			number /= i
			limit = int64(math.Sqrt(float64(number)))
		}
	}
	fmt.Printf("Наибольший простой делитель %d\n", number)
}
