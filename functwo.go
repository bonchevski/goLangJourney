package main

import "fmt"

//  the _ ignores the second argument
func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome", firstName)
}

func getNames() (string, string) {
	return "Foo", "Bar"
}
