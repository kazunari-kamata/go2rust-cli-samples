package main

import "fmt"

func countdown(start int) {
	count := start
	for count > 0 {
		fmt.Println(count)
		count = count - 1
	}
}

func tickOnce() {
	for {
		fmt.Println("tick")
		return
	}
}

func main() {
	countdown(3)
	tickOnce()
}
