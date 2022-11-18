package main

import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

//struct method / struct function
func (customer Customer) sayHai(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuu from", a.Name)
}

func main() {
	var fauzi Customer
	fauzi.Name = "Achmad Sidik Fauzi"
	fauzi.Address = "Indonesia"
	fauzi.Age = 23

	fmt.Println(fauzi)

	//cara yang sama 
	sidik := Customer{
		Name: "Adelia Risky Santoso",
		Address: "Indonesia",
		Age: 21,
	}

	fmt.Println(sidik)

	achmad := Customer{"Achmad Risky Fadelia", "Indonesia", 100}
	fmt.Println(achmad)

	//call struct method / struct function
	fauzi.sayHai("Fasi")
	fauzi.sayHuuu()
}