package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan ke ", i)
	}

	slice := []string{"Achmad", "Sidik", "Fauzi", "Adelia", "Risky", "Santoso"}
	for i:=0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//cara cepat ambuil semua data For Range
	for i, value := range slice {
		fmt.Println("Index ", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Fauzi"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
