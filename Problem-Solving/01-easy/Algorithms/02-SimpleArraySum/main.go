package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var result int32 = 0
	for _, data := range ar {
		result += data
	}
	return result
}

func main() {
	var ar = [...]int32{1, 2, 3, 4, 10, 11}
	// karena parameter fungsinya slice, kita convert dulu dari array ke slice
	// ambil semua elemen array dengan slice [:]
	result := simpleArraySum(ar[:])
	fmt.Println(result)
}
