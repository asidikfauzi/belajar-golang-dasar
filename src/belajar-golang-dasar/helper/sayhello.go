package helper

import "fmt"

var mantap = "uhuy"
var Application = "1.1.0"

func SayHello(name string){
	fmt.Println("Hello", name)
}

//tidak bisa di akses oleh file lain jika nama fungsinya diawali dengan huruf kecil
func sayGoodBye(name string){
	fmt.Println("Goodbye", name)
}