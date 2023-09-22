package main

import (
	"fmt"

	"github.com/santos95mat/go-utils/CPF"
)

func main() {
	var cpf string = "054.040.061-05"
	isValid, err := CPF.IsValid(cpf)

	fmt.Println(cpf, isValid, err)
}
