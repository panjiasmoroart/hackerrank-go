package main

import (
	illustration "05-DiagonalDifference/illustration"
	"fmt"
)

/*
// Ilustrasi input
arr := [][]int32{
    {11, 2, 4},
    {4, 5, 6},
    {10, 8, -12},
}

Primary diagonal (kiri ke kanan)
+-----+-----+------+
| +11 |   2 |    4 | // baris ke 0
+-----+-----+------+
|   4 | +5  |    6 | // baris ke 1
+-----+-----+------+
|  10 |   8 | -12  | // baris ke 2
+-----+-----+------+
Primary diagonal = 11 + 5 + (-12) = 4

// table ilustrasi arr[row][column]
// arr[i][i]
| i | Elemen    | Nilai |
| - | --------- | ----- |
| 0 | arr[0][0] | 11    |
| 1 | arr[1][1] | 5     |
| 2 | arr[2][2] | -12   |

//Secondary diagonal (kanan ke kiri)
+-----+-----+------+
|  11 |   2 |  +4  | // baris ke 0
+-----+-----+------+
|   4 |  +5 |    6 | // baris ke 1
+-----+-----+------+
| +10 |   8 |  -12 | // baris ke 2
+-----+-----+------+
Secondary diagonal = 4 + 5 + 10 = 19

// table ilustrasi arr[row][column]
arr[len(arr) - 1 - i][i]
| i | Elemen    | Rumus         | Nilai |
| - | --------- | ------------- | ----- |
| 0 | arr[2][0] | arr[3-1-0][0] | 10    |
| 1 | arr[1][1] | arr[3-1-1][1] | 5     |
| 2 | arr[0][2] | arr[3-1-2][2] | 4     |
*/

func main() {
	arr := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	illustration.DiagonalDifferenceShow(arr)

	result := diagonalDifference(arr)
	fmt.Println("Hasilnya = ", result)
}

func diagonalDifference(arr [][]int32) int32 {
	var sumPrimary, sumSecondary int32
	for i := 0; i < len(arr); i++ {
		// Diagonal kiri ke kanan (primary diagonal)
		/*
		   | i | Elemen    | Nilai |
		   | - | --------- | ----- |
		   | 0 | arr[0][0] | 11    |
		   | 1 | arr[1][1] | 5     |
		   | 2 | arr[2][2] | -12   |
		*/
		sumPrimary += arr[i][i]

		// diagonal kanan ke kiri (secondary diagonal)
		// ambil elemen arr[len(arr)-1-i][i]
		// dengan panjang array = 3:
		// arr[row][column]
		/*
		   | i | Elemen    | Rumus         | Nilai |
		   | - | --------- | ------------- | ----- |
		   | 0 | arr[2][0] | arr[3-1-0][0] | 10    |
		   | 1 | arr[1][1] | arr[3-1-1][1] | 5     |
		   | 2 | arr[0][2] | arr[3-1-2][2] | 4     |
		*/
		sumSecondary += arr[len(arr)-1-i][i]
	}

	/*
	   sumPrimary = 4
	   sumSecondary = 19
	   diff = 4 - 19 = -15  (negatif)
	   if diff < 0 {
	       diff = -diff     // -(-15) = 15
	   }
	*/

	diff := sumPrimary - sumSecondary
	// jika hasil perhitungan negatif
	if diff < 0 {
		// balikkan tanda negatifnya menjadi positif
		diff = -diff
	}
	return diff
}
