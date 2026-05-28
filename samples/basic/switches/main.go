package main

import "fmt"

func describe(count int) {
	switch count {
	case 0:
		fmt.Println("zero")
	case 1, 2:
		fmt.Println("small")
	default:
		fmt.Println("many")
	}
}

func main() {
	describe(2)
}
