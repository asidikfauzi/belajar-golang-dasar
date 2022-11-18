package main

import "fmt"

func sumAll(numbers ...int)int {
	total := 0

	for _, values := range numbers {
		total += values
	}
	return total
}

func main() {
	total := sumAll(10,90,20,30)
	fmt.Println(total)
	

	//slice parameter
	slice := []int{18,10,19,60,11}
	total = sumAll(slice...)
	fmt.Println(total)

}
