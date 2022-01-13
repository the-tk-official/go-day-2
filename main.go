package main

import "log"

func main() {
	var isTrue bool

	isTrue = false

	if !isTrue {
		log.Println("False")
	} else {
		log.Println("True")
	}

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("My variable is cat.")

	case "dog":
		log.Println("My variable is dog.")

	case "fish":
		log.Println("My variable is fish.")

	default:
		log.Println("My variable is None")

	}
}
