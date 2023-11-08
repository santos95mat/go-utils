package main

import (
	"fmt"

	passwordGenerator "github.com/santos95mat/go-utils/PasswordGenerator"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(passwordGenerator.Generate(12))
	}
}
