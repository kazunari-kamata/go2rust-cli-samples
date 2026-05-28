package main

import "fmt"

func describe(count int) {
	switch {
	case count > 10:
		fmt.Println("large")
	case count > 0, count == -1:
		fmt.Println("known")
	default:
		fmt.Println("zero")
	}
}

func main() {
	describe(3)
}
