package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func main() {
	c := sum(2, 3)
	fmt.Println(c)
}