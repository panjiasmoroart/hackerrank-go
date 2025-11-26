package main

import "fmt"

func main() {
	arr := []int32{73, 67, 38, 33}
	results := gradingStudents(arr)
	fmt.Println(results)
}

/*
   grade = 84 round to 85 (85 - 84 is less than 3)
   grade = 29 do not round (results is less than 38)
   grade = 57 do not round (60 - 57 is 3 or higher)
*/

func gradingStudents(grades []int32) []int32 {
	results := make([]int32, len(grades))

	for index, grade := range grades {
		// if the value of  is less than , no rounding occurs as the result will still be a failing grade (nilai gagal)
		// Nilai di bawah 38 tidak dibulatkan
		if grade < 38 {
			results[index] = grade
		} else {
			// If the difference (selisih) between the  and the next multiple of 5  is less than 3 , round  up to the next multiple of 5

			// Loop tiap nilai
			// Nilai 1: 73
			// (73 / 5 = 14) (14 + 1 = 15) 15 * 5 = 75
			// Selisih = 75 - 73 = 2
			// Karena 2 < 3, dibulatkan ke 75
			// results = [75, 0, 0, 0]

			// Nilai 2: 67
			// (67 / 5 = 13) (13 + 1 = 14)  14 * 5 = 70
			// Selisih = 70 - 67 = 3
			// Karena 3 tidak < 3, tidak dibulatkan
			// results = [75, 67, 0, 0]

			// Nilai 3: 38
			// (38 / 5 = 7) (7 + 1 = 8) 8 * 5 = 40
			// Selisih = 40 - 38 = 2
			// Karena 2 < 3, dibulatkan ke 40

			// Nilai 4: 33
			// 33 < 38  langsung tidak dibulatkan
			// results = [75, 67, 40, 33]

			nextMultiple := ((grade / 5) + 1) * 5
			// // Cari kelipatan 5 terdekat berikutnya

			if nextMultiple-grade < 3 {
				results[index] = nextMultiple
			} else {
				results[index] = grade
			}
		}
	}

	return results
}

func gradingStudentsPrintSlice(grades []int32) []int32 {
	for _, value := range grades {
		fmt.Printf("%d ", value)
	}

	return grades
}
