package main

import "fmt"

func main() {
	for i := 1; i < 100000; i++ {
		if isPerfectNumber(i) {
			fmt.Println(i)
		}
	}
}

func isPerfectNumber(i int) bool {
	divisors := []int{}
	result := 0

	if i%2 != 0 {
		return false
	}

	// Technically, we can stop at i / 2,
	// but the Python version has the same bug
	for x := 1; x < i; x++ {
		if i%x == 0 {
			divisors = append(divisors, x)
		}
	}

	for _, x := range divisors {
		result += x
	}

	if result == i {
		return true
	}

	return false
}
