package main

import "fmt"

func sayHello(firstname string, lastName string) {
	fmt.Println("Hello ", firstname, lastName)
}

func main() {
	sayHello("Fauzi", "Yarro");
}
