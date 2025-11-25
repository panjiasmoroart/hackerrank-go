package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputPm := "12:01:00PM"
	// inputAm := "12:01:00AM"
	result := timeConversion(inputPm)
	fmt.Println(result)
}

// fungsi untuk mengubah format waktu 12-hour (AM/PM) menjadi format 24-hour (military time)
func timeConversion(s string) string {
	// s := "12:01:00PM"
	// return "12:01:00"
	// s := "12:01:00AM"
	// return "00:01:00"
	s = strings.ToLower(s)
	// cek apakah string str diakhiri oleh suffix tertentu.
	// strings.HasSuffix("hello.mp3", ".mp3") // true
	// strings.HasSuffix("hello.mp3", ".txt") // false
	pm := strings.HasSuffix(s, "pm")
	am := strings.HasSuffix(s, "am")

	// 12:01:00
	t := s[:len(s)-2]

	timerArr := strings.Split(t, ":")
	hh, mm, ss := timerArr[0], timerArr[1], timerArr[2]
	hhInt, err := strconv.Atoi(hh) // convert string to int untuk operasi matematichs

	if err != nil {
		panic(err.Error())
	}

	// conversi 12hour ke 24hour
	// 12 AM = 00 (tengah malam)
	// Selain 12, AM tetap sama (misal 08 AM tetap 08)
	if am && hhInt == 12 {
		hhInt = 0
	}

	// 01 PM ke 13
	// 05 PM ke 17
	// 12 PM tetap 12 (tidak ditambah)
	if pm && hhInt != 12 {
		hhInt += 12
	}

	hh = strconv.Itoa(hhInt) // int to string

	return fmt.Sprintf("%02s:%02s:%02s", hh, mm, ss)
}

func timeConversionExample(s string) string {
	// s := "12:01:00PM"
	// return "12:01:00"
	// s := "12:01:00AM"
	// return "00:01:00"
	s = strings.ToLower(s)
	// pm := strings.HasSuffix(s, "pm")
	// am := strings.HasSuffix(s, "am")

	// 12:01:00
	t := s[:len(s)-2]

	return t
}
