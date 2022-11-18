package main

import "fmt"

func main() {
	name := "Fauzi"
	if name == "Fauzi" {
		fmt.Println("Hallo Fauzi Yarro")
	} else if name == "Adelia" {
		fmt.Println("Yah ini lah aku")
	} else {
		fmt.Println("Kamu siapa anying ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
	
}
