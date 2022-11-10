package my_error

import (
	"fmt"
	"errors"
)

func SayHi(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name is empty")
	}

	message := fmt.Sprintf("Hi %v. Welcome!", name)
	return message, nil 
}