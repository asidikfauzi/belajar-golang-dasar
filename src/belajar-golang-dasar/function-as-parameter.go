package main

import "fmt"

type Filter func(string)string

//bisa
// type Filter func(string, string, string)string

func sayHelloWithFilter(name string, filter Filter)  {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	//sama aja
	sayHelloWithFilter("Fauzi", spamFilter)

	result := spamFilter
	sayHelloWithFilter("Anjing", result)
}