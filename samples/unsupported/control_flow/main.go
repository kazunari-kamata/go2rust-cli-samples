package main

import "fmt"

func main() {
	var count int = 3
	if count > 0 {
		fmt.Println("positive")
	}
	for count > 0 {
		fmt.Println("loop")
		count = count - 1
	}
}
