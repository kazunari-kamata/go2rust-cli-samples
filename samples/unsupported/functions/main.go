package main

import "fmt"

type counter struct {
	value int
}

func (c counter) valueOrZero() int {
	return c.value
}

func main() {
	var count int = 1
	item := counter{value: count}
	fmt.Println(item.valueOrZero())
	fmt.Println(count)
}
