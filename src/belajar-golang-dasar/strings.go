package main

import (
	"fmt"
	"strings"
)

func main(){
	//apakah dalam string ini ada kata Fauzi?
	fmt.Println(strings.Contains("Achmad Sidik Fauzi", "Fauzi")) //true

	fmt.Println(strings.Split("Achmad Sidik Fauzi", " ")) //menjadi slice

	fmt.Println(strings.ToLower("Achmad Sidik Fauzi")) //menjadi lower case
	fmt.Println(strings.ToUpper("Achmad Sidik Fauzi")) //menjadi upper case
	fmt.Println(strings.ToTitle("achmad sidik fauzi")) //menjadi upper case
	fmt.Println(strings.Trim("         Achmad Sidik Fauzi     ", " ")) //kiri kanan ilang
	fmt.Println(strings.ReplaceAll("Eko Yarro Eko", "Eko", "Budi")) //replace Eko jadi Budi
}