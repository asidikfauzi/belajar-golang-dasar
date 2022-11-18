package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("epo"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eko eka epo edo eto eyo eki", 10)
	fmt.Println(search)
}
