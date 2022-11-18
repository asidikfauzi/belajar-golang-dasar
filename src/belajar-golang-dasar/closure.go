package main

import "fmt"


func main() {
	counter := 0
	name := "Fauzi"

	increment := func() {
		name := "Adelia"
		fmt.Println(name)
		counter++
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}