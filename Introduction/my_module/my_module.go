package my_module

import "fmt"

func SayHi(name string) string {
	message := fmt.Sprintf("Hi %v. Welcome!", name)
	return message
}