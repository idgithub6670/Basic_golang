package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Sunday

	switch today {
	case time.Monday:
		fmt.Println("Today is Monday")
	case time.Tuesday:
		fmt.Println("Today is Tuesday")
	case time.Wednesday:
		fmt.Println("Today is Wednesday")
	case time.Thursday:
		fmt.Println("Today is Thursday")
	case time.Friday:
		fmt.Println("Today is Friday")
	case time.Saturday:
		fmt.Println("Today is Saturday")
	case time.Sunday:
		fmt.Println("Today is Sunday")

	}
}
