package main

import (
	"fmt"

	"Go-REST-API/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
