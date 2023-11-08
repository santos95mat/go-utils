package CNPJ

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

func Generate() string {
	randon := rand.Intn(99999999-10000000) + 10000000
	randonSTR := strconv.Itoa(randon) + "0001"

	return newCNPJ(randonSTR)
}

func IsValid(cnpj string) (bool, string) {
	var cnpjIsValid bool = false

	formatCNPJ, err := formatCNPJ(cnpj)

	if err != nil {
		return cnpjIsValid, cnpj
	}

	newCNPJ := newCNPJ(formatCNPJ[:12])

	if newCNPJ == formatCNPJ {
		cnpjIsValid = true
	}

	return cnpjIsValid, newCNPJ
}

func formatCNPJ(cnpj string) (string, error) {
	var err error = nil
	var format string = cnpj

	format = strings.Replace(format, ".", "", -1)
	format = strings.Replace(format, "-", "", -1)
	format = strings.Replace(format, "/", "", -1)

	if len(format) != 14 {
		err = errors.New("error: Formato de CNPJ inválido")
		return format, err
	}

	_, err = strconv.Atoi(format)

	if err != nil {
		err = errors.New("error: CNPJ deve conter apenas [numeros - . /]")
		return format, err
	}

	return format, nil
}

func newCNPJ(cnpj string) string {
	var newCNPJ string

	newCNPJ = cnpj + strconv.Itoa(generateDigito(cnpj))
	newCNPJ = newCNPJ + strconv.Itoa(generateDigito(newCNPJ))

	return newCNPJ
}

func generateDigito(cnpj string) int {
	var i int = 9
	var s int = 0

	for x := len(cnpj) - 1; x >= 0; x-- {
		if i == 1 {
			i = 9
		}

		y, _ := strconv.Atoi(string(cnpj[x]))
		s = s + (y * i)
		i--
	}

	r := s % 11

	return r
}
