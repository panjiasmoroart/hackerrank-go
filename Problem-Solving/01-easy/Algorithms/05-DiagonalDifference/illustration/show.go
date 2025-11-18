package illustration

import "fmt"

func DiagonalDifferenceShow(arr [][]int32) {
	// Loop untuk diagonal 1 (kiri-atas → kanan-bawah)
	fmt.Println("Diagonal 1 (kiri-atas → kanan-bawah):")
	for i := 0; i < len(arr); i++ {
		value := arr[i][i]
		fmt.Printf("Iterasi %d: arr[%d][%d] = %d\n", i, i, i, value)
	}
	fmt.Println("Primary diagonal = 11 + 5 + (-12) = 4")
	fmt.Println()

	// Loop untuk diagonal 2 (kanan-atas → kiri-bawah)
	fmt.Println("Diagonal 2 (kanan-atas → kiri-bawah):")
	for i := 0; i < len(arr); i++ {
		value := arr[len(arr)-1-i][i]
		fmt.Printf("Iterasi %d: arr[%d][%d] = %d\n", i, len(arr)-1-i, i, value)
	}
	fmt.Println("Secondary diagonal = 4 + 5 + 10 = 19")
	fmt.Println(" diff = 4 - 19 = -15  (negatif)")
	fmt.Println()
}
