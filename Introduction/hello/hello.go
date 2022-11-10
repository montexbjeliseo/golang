package main

import (
	"fmt"
	"example.com/my_module"
)

func main(){
	message := my_module.SayHi("Eliseo")
	fmt.Println(message)
}