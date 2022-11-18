package main

import(
	"fmt"
	"container/list"
)

func main() {
	data := list.New()
	data.PushBack("Kurniawan") //pushback kirim data paling ujung
	data.PushBack("Fauzi")
	data.PushFront("Asidik") // pushfront kirim data paling depan

	// data.Front().Next().Next() //ini artinya kirim data yang nextnya 2 kali dari depan
	// data.Front().Prev().Prev() //ini artinya kirim data yang sebelumnya 2 kali dari depan tapi hasilnya nil

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	//dari depan ke belakang
	// for element:= data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	// dari belakang ke depan
	for element:= data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}