package main

import "fmt"

func sampleInput() {
	ar := []int64{
		1000000001,
		1000000002,
		1000000003,
		1000000004,
		1000000005,
	}

	result := aVeryBigSum(ar)
	fmt.Println("Hasil penjumlahan:", result)
}

func main() {
	sampleInput()
}

func aVeryBigSum(ar []int64) int64 {
	var sum int64 = 0
	// Write your code here
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum
}
