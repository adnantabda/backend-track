package main

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(Sum([]int{}))
}
