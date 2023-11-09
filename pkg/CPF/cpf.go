package CPF

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

func Generate() string {
	numbers := "0123456789"
	randonStr := ""
	for i := 0; i < 9; i++ {
		randon := rand.Intn(10)
		randonStr += string(numbers[randon])
	}

	newCPF := newCPF(randonStr)
	newCPF = FomartCPF(newCPF)

	return newCPF
}

func IsValid(cpf string) (bool, string) {
	var cpfIsValid bool = false

	cpf, err := isValidFormat(cpf)

	if err != nil {
		return cpfIsValid, cpf
	}

	if isSequency(cpf) {
		return cpfIsValid, cpf
	}

	newCPF := newCPF(cpf[:9])

	if newCPF == cpf {
		cpfIsValid = true
	}

	newCPF = FomartCPF(newCPF)

	return cpfIsValid, newCPF
}

func FomartCPF(cpf string) string {
	cpf = cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
	return cpf
}

func isValidFormat(cpf string) (string, error) {
	var err error = nil
	var format string = cpf

	format = strings.Replace(format, ".", "", -1)
	format = strings.Replace(format, "-", "", -1)

	if len(format) != 11 {
		err = errors.New("error: Formato de CPF inválido")
		return format, err
	}

	_, err = strconv.Atoi(format)

	if err != nil {
		err = errors.New("error: CPF deve conter apenas [numeros - .]")
		return format, err
	}

	return format, err
}

func isSequency(cpf string) bool {
	var r = rune(cpf[0])

	for _, l := range cpf {
		if r != l {
			return false
		}
	}

	return true
}

func newCPF(cpf string) string {
	var newCPF string

	newCPF = cpf + strconv.Itoa(generateDigito(cpf))
	newCPF = newCPF + strconv.Itoa(generateDigito(newCPF[1:]))

	return newCPF
}

func generateDigito(cpf string) int {
	var i int = 10
	var s int = 0
	var r int = 0
	var d int = 0

	for _, x := range cpf {
		y, _ := strconv.Atoi(string(x))
		s = s + (y * i)
		i = i - 1
	}

	r = s % 11

	if r < 2 {
		d = 0
	} else {
		d = 11 - r
	}

	return d
}
