package main

import "fmt"

func sayHello(name string) string {
	if name == "" {
		return "Hallo Bro"
	}else{
		return "Hallo " + name
	}
}

func main() {
	result := sayHello("Fauzi")
	fmt.Println(result)
}
