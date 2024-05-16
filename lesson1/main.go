package main

import (
	"fmt"
	_ "fmt"
	"lesson1/greet"
)

func main() {
	var name string
	var table int

	fmt.Print("Enter name:")
	fmt.Scanln(&name)

	fmt.Print("Enter table number:")
	fmt.Scanln(&table)

	text := greet.Greet(name, table)
	fmt.Println(text)
}
