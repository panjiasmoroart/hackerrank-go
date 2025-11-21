package main

import "fmt"

func main() {
	arr := []int32{1, 3, 5, 7, 9}
	miniMaxSum(arr)
}

func miniMaxSum(arr []int32) {
	var total int64 = 0
	// Write your code here
	for _, value := range arr {
		total += int64(value)
	}

	// cari min dan max
	minValue := arr[0]
	maxValue := arr[0]
	for _, v := range arr {
		if v < minValue {
			minValue = v
		}
		if v > maxValue {
			maxValue = v
		}
	}

	// hitung sum nya
	// total 25, minnya 1 dan maxnya 9
	// 25 - 9 = 16 sama dengan menjumlahkan [1, 3, 5, 7]
	minSum := total - int64(maxValue)
	// 25 - 1 = 24 sama dengan menjumlahkan [3, 5, 7, 9]
	maxSum := total - int64(minValue)

	fmt.Println(minSum, maxSum)
}
