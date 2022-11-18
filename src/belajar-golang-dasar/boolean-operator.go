package main

import "fmt"

func main() {
	var ujian = 83
	var absensi = 83

	var lulusUjian = ujian > 81
	var lulusAbsensi = absensi > 81

	var lulus = lulusAbsensi && lulusUjian

	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)

}