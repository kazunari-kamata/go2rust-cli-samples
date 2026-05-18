package main

import "fmt"

func add(left, right int) int {
	total := left + right
	return total
}

func main() {
	fmt.Println(add(1, 2))
}
