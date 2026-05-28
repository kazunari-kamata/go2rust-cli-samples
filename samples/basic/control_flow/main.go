package main

import "fmt"

func main() {
	count := 3
	if count > 10 {
		fmt.Println("large")
	} else if count > 0 {
		fmt.Print("positive")
		fmt.Println(" count")
	} else {
		fmt.Println("zero")
	}
}
