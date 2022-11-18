package main

import "fmt"

func getFullName()(string, string) {
	return "Fauzi", "Yarro"
}

func main() {
	// firstname, lastName := getFullName();
	// fmt.Println(firstname)
	// fmt.Println(lastName)

	//jika tidak peduli dengan lastname menggunakan garis bawah
	firstname, _ := getFullName();
	fmt.Println(firstname)
}
