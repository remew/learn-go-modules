package main

import (
	"fmt"
	"learn-go-modules/lib"
)

func main() {

	var name string
	fmt.Print("name:")
	fmt.Scan(&name)

	var age int
	fmt.Print("age :")
	fmt.Scan(&age)

	person := lib.Person{name, age}
	fmt.Printf("Hello %s. you are %d years old.\n", person.Name, person.Age)
}
