package main

import "fmt"

func main() {
	const firstName string = "Achmad Sidik"
	const lastName = "Fauzi"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
	const (
		firstName2 string = "Achmad Sidik"
		lastName2         = "Fauzi"
		value2            = 1000
	)

	fmt.Println(firstName2)
	fmt.Println(lastName2)
	fmt.Println(value2)

}
