package main

import "fmt"

/*
// sample input
arr := []int32{-4, 3, -9, 0, 4, 1}
plusMinus(arr)

hitung jumlah bilangan positif, negatif, dan nol dalam sebuah array
kemudian kita bagi dengan banyak array yang di input
*/

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}

func plusMinus(arr []int32) {
	positive, negative, zero := 0, 0, 0

	// Write your code here
	for _, value := range arr {
		if value > 0 {
			positive++
		} else if value < 0 {
			negative++
		} else {
			zero++
		}
	}

	countArr := float64(len(arr))

	fmt.Printf("%.6f\n", float64(positive)/countArr)
	fmt.Printf("%.6f\n", float64(negative)/countArr)
	fmt.Printf("%.6f\n", float64(zero)/countArr)
}
