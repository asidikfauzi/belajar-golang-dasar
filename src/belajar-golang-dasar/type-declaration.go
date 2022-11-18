package main

import "fmt"

func main(){
	type NoKTP string
	type Married bool
	var NoKtpFauzi NoKTP = 	"3520100607990004";
	var marriedStatus Married= true;
	fmt.Println(NoKtpFauzi)
	fmt.Println(marriedStatus)
}