package main

import "fmt"

type Blacklist func(string)bool

//bisa
// type Filter func(string, string, string)string

func registerUser(name string, blakckList Blacklist)  {
	if blakckList(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blackListAdmin(name string) bool {
// 	return name == "admin"
// }

// func blackListRoot(name string) bool {
// 	return name == "root"
// }

func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blackList)
	registerUser("eko", blackList)

	//sama saja
	registerUser("eko", func(name string) bool {
		return name == "root"
	})
}