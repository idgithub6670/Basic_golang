package main

import "fmt"

func main() {
	status := map[int]string{
		200: "OK",
		400: "Bad Request",
		404: "Not found",
		500: "Internal Server Error",
	}
	fmt.Println(status)

	status[200] = "Oke"
	fmt.Println(status)

	delete(status, 200)
	fmt.Println(status)
}
