package greetings

import "fmt"

/*
This function takes a name parameter whose type is string. The function also returns a string. In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name. For more about exported names, see Exported names in the Go tour.
*/

/*
The long way.
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
*/

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
