package main

import "fmt"

func main() {
	var langs = []string{"Golang", "Java", "Python", "C#"}
	fmt.Println(langs)
	fmt.Println(langs[0])

	langs[0] = "HTML"
	fmt.Println(langs)
	fmt.Println(langs[0])
}
