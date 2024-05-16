package main

import "fmt"

type User struct {
	UserName string
	FullNmae string
	Age      int
}

func main() {
	username := "Thana"
	fullname := "Thana Praserttha"
	age := 23

	u := User{
		UserName: username,
		FullNmae: fullname,
		Age:      age,
	}
	fmt.Print(u)
}
