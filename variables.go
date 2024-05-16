package main

import "fmt"

func main() {
	var name string = "Thana"
	age := 23

	fmt.Println(name, " Age:", age)
	fmt.Printf("My name is %s i'm %d years old \n", name, age)

	const (
		zero = iota
		one
		two
		three
		four
		five
	)
	fmt.Println(one, two)
}
