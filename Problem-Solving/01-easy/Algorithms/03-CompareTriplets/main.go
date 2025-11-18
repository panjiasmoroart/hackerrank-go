package main

import "fmt"

// ilustrasi
/*
a = [1, 2, 3]  (Alice)
b = [3, 2, 1]  (Bob)

// table perbandingan
| Kategori       | Alice (a[i]) | Bob (b[i]) | Pemenang      |
| -------------- | ------------ | ---------- | ------------- |
| Clarity (0)    | 1            | 3          | Bob           |
| Originality(1) | 2            | 2          | Seri (0 poin) |
| Difficulty (2) | 3            | 1          | Alice         |

// Langkah 2: Hitung skor
Alice menang 1 kategori → Alice = 1
Bob menang 1 kategori → Bob = 1
Seri 1 kategori → tidak ada poin

Hasil akhir: [1, 1]
*/

func main() {
	a := []int32{1, 2, 3} // (Alice)
	b := []int32{3, 2, 1} // (Bob)
	result := compareTriplets(a, b)
	fmt.Println(result)
}

func compareTriplets(a []int32, b []int32) []int32 {
	var aliceScore, bobScore int32

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			aliceScore++
		} else if a[i] < b[i] {
			bobScore++
		}
	}

	return []int32{aliceScore, bobScore}
}
