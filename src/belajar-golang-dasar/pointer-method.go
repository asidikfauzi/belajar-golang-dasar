package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	fauzi := Man{"Fauzi"}
	fauzi.Married()
	fmt.Println(fauzi.Name)
}