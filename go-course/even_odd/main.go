package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for n := range nums {
		phrase := fmt.Sprint(n) + " is " + evenOrOdd(n)
		fmt.Println(phrase)
	}
}

func evenOrOdd(n int) string {
	if n%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}
