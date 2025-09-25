package main

import (
	"fmt"
)

func main() {
	var name string = "John"
	fmt.Printf("Hello, %s\n", name)
	age := 24
	fmt.Printf("my age is %d\n", age)
	var city string
	city = "Dallas"
	fmt.Printf("I am in %s\n", city)
	var country, continent string = "USA", "North America"
	fmt.Println(country, continent)
}
