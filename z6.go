package main

func main() {
	sum_k := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		sum_k += i * i
		sum += i
	}
	square_sum := sum * sum
	difference := sum_k - square_sum
	println("Difference: ", difference)
}
