package main

import (
	"fmt"
	"example.com/my_error"
	"log"
)

func main(){
	log.SetPrefix("my_error: ")
	log.SetFlags(0)

	message, err := my_error.SayHi("")
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}