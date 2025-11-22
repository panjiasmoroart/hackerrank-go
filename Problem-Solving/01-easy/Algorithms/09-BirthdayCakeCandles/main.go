package main

import "fmt"

func main() {
	arr := []int32{4, 4, 1, 3}
	result := birthdayCakeCandles(arr)
	fmt.Println(result)
}

func birthdayCakeCandles(candles []int32) int32 {
	var max int32 = 0
	var count int32 = 0

	// Write your code here
	for _, value := range candles {
		if value > max {
			max = value
			count = 1
		} else if value == max {
			count++
		}
	}

	return count
}

/*
arr := []int32{4, 4, 1, 3}

Iterasi 1
value = 4
4 > 0 benar

update:
max = 4
count = 1   // lilin tinggi 4 ditemukan 1x

Iterasi 2
value = 4
4 > max (4) salah
4 == max (4) benar
increment:
count = 2   // ada lilin tinggi 4 lagi
*/

func birthdayCakeCandlesExam(candles []int32) int32 {
	var arrCandles int32
	for _, value := range candles {
		fmt.Print(value)
		arrCandles = value
	}

	return arrCandles

}
