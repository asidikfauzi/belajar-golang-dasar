package main

import "fmt"

type HasName interface {
	GetName() string 
}


func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var fauzi Person
	fauzi.Name = "Fauzi"
	sayHello(fauzi)

	cat := Animal {
		Name: "Luffy",
	}
	sayHello(cat)
}