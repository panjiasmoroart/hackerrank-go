package main

import "fmt"

func main() {
	var a, b uint32 = 3, 7
	result := solveMeFirst(a, b)
	fmt.Println(result)
}

func solveMeFirst(a uint32, b uint32) uint32 {
	// Hint: Type return (a+b) below
	return a + b
}
