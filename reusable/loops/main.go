package main

import "fmt"

func rng(start, end int) []int {
	result := []int{}

	for i := start; i <= end; i += 1 {
		result = append(result, i)
	}

	return result
}

func rngOdd(start, end int) []int {
	result := []int{}

	for i := start; i <= end; i += 1 {
		if i%2 != 0 {
			result = append(result, i)
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	rng(15, 30)
	rngOdd(15, 30)
}
