package main

import "fmt"

func endApp(){
	fmt.Println("Aplikasi Selesai")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message : ", message)
	}
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi ERROR")
	}
	
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Fauzay")
}