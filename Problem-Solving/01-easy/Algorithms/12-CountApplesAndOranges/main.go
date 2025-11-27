package main

import "fmt"

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */

/*
1. Input:
s dan t: posisi awal dan akhir rentang rumah.
a: posisi pohon apel.
b: posisi pohon jeruk.
apples: daftar jarak buah apel dari pohonnya.
oranges: daftar jarak buah jeruk dari pohonnya.

Output: Dua angka yang dipisahkan oleh baris baru, yang mewakili jumlah apel dan jeruk yang jatuh di rentang rumah.
Logika: Sebuah apel atau jeruk dianggap jatuh di rentang rumah jika posisinya (pohon + jarak) berada di antara s dan t (inklusif).

2. Contoh :
s = 7, t = 10
a = 4
b = 12
apples = [2, 3, -4]
oranges = [3, -2, -4]

3. Hitung posisi apel a = 4
Apel 1: 4 + 2  = 6 (Tidak termasuk dalam rentang 7 hingga 10)
Apel 2: 4 + 3  = 7 (Termasuk dalam rentang)
Apel 3: 4 + -4 = 0 (Tidak termasuk dalam rentang 7 hingga 10)
Jumlah apel yang jatuh di rentang rumah adalah 1

4. Hitung posisi jeruk b = 12
Jeruk 1: 12 + 3    = 15 (Tidak termasuk dalam rentang)
Jeruk 2: 12 + (-2) = 10 (Termasuk dalam rentang)
Jeruk 3: 12 + (-4) = 8 (Termasuk dalam rentang)
Jumlah jeruk yang jatuh di rentang rumah adalah 2
*/

func main() {
	arrApples := []int32{2, 3, -4}
	arrOrange := []int32{3, -2, -4}
	var s int32 = 7
	var t int32 = 10
	var a int32 = 4
	var b int32 = 12

	countApplesAndOranges(s, t, a, b, arrApples, arrOrange)
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var applesCount int32
	var orangesCount int32

	// Write your code here
	// Hitung apel
	for _, apple := range apples {
		if a+apple >= s && a+apple <= t {
			applesCount++
		}
	}

	// Hitung jeruk
	for _, orange := range oranges {
		if b+orange >= s && b+orange <= t {
			orangesCount++
		}
	}

	fmt.Println(applesCount)
	fmt.Println(orangesCount)
}
