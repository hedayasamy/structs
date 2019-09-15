package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//zero value struct declaration
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
