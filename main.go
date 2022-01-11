package main

import "log"

func main() {
	var myString = "Black"
	log.Println("My string is set to", myString)

	changeUsingPointer(&myString)
	log.Println("After func my string is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}
