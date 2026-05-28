package main

import "fmt"

func logIfEnabled(enabled bool) {
	if !enabled {
		return
	}
	fmt.Println("enabled")
}

func main() {
	logIfEnabled(true)
}
