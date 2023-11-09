package CNPJ

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

func Generate() string {
	numbers := "0123456789"
	randonStr := ""
	for i := 0; i < 8; i++ {
		randon := rand.Intn(10)
		randonStr += string(numbers[randon])
	}
	randonStr += "0001"

	cnpj := newCNPJ(randonStr)
	cnpj = formatCNPJ(cnpj)
	return cnpj
}

func IsValid(cnpj string) (bool, string) {
	var cnpjIsValid bool = false

	cnpj, err := isValidFormat(cnpj)

	if err != nil {
		return cnpjIsValid, cnpj
	}

	newCNPJ := newCNPJ(cnpj[:12])

	if newCNPJ == cnpj {
		cnpjIsValid = true
	}

	newCNPJ = formatCNPJ(newCNPJ)

	return cnpjIsValid, newCNPJ
}

func formatCNPJ(cnpj string) string {
	cnpj = cnpj[:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:]

	return cnpj
}

func isValidFormat(cnpj string) (string, error) {
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
	if r == 10 {
		return 0
	}

	return r
}
