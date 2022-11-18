package main

import "fmt"

func main() {
	name := "Fauzi"
	
	switch name {
		case "Fauzi":
			fmt.Println("Nama dah bener")
			fmt.Println("Nama dah bener")
		case "Adelia" :
			fmt.Println("Ini nama siapa")
			fmt.Println("Ini nama siapa")
		default :
			fmt.Println("Cari lagi nama kamu")
			fmt.Println("Cari lagi nama kamu")
	}

	//switch dengan ekspresi
	// switch length := len(name); length > 5 {
	// 	case true:
	// 		fmt.Println("Nama Terlalu Panjang")
	// 	case false :
	// 		fmt.Println("Nama Sudah Benar")
	// }

	//switch tanpa ekspresi
	length := len(name); 
	switch {
		case length > 10:
			fmt.Println("Nama Terlalu Panjang")
		case length > 5:
			fmt.Println("Nama Lumayan Benar")
		default : 
			fmt.Println("Nama Sudah Benar")
	}
}
