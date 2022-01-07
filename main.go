package main

import "log"

func main() {
	log.Println(saySomething("Hello, "))
	log.Println(saySomething("Goodbye, "))

	var whatToSay string
	var saySomethingElse string
	var i int

	whatToSay, _ = saySomething("Tamerlan!")

	log.Println(whatToSay)

	saySomethingElse, _ = saySomething("King!")

	log.Println(saySomethingElse)

	i = 1
	i = 10
	log.Println(i)
}

func saySomething(s string) (string, string) {
	return s, "World!"
}
