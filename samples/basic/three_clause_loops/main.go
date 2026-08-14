package main

import "fmt"

func countTo(limit int) {
	for count := 0; count < limit; count++ {
		fmt.Println(count)
	}
}

func countdown() {
	count := 3
	for ; count > 0; count-- {
		fmt.Println(count)
	}
}

func main() {
	countTo(3)
	countdown()
}
