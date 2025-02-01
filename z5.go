package main

import "fmt"

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	result := 1
	for i := 1; i <= 20; i++ {
		result = lcm(result, i)
	}
	fmt.Println("Наименьший общий кратный", result)
}
