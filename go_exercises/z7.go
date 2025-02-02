package main

func is_Prime(a int) bool {
	if a < 2 {
		return false
	}
	for i := 2; i*i <= a; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	count := 0
	for i := 2; ; i++ {
		if is_Prime(i) {
			count++
		}
		if count == 10001 {
			println("1001-ое простое число", i)
			break
		}
	}

}
