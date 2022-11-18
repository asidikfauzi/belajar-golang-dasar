package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Achmad"
	names[1] = "Sidik"
	names[2] =  "Fauzi"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}

	var data = [10]string {
		
	}

	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))
	fmt.Println(len(data))
}
