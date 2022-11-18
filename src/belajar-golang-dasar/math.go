package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println(math.Round(1.7)) // pembulatan > 5 keatas
	fmt.Println(math.Round(1.3)) // < 5 kebawah

	fmt.Println(math.Floor(1.7)) // pembulatan ke bawah
	fmt.Println(math.Ceil(1.3)) // pembulatan ke keatas
	fmt.Println(math.Max(1,7)) // nilai Max
	fmt.Println(math.Min(1,7)) // nilai Min




}