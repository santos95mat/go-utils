package password

import (
	"math/rand"
)

var (
	//Lower Case
	lowCase = "abcdefghijklmnopqrstuvwxyz"
	//Upper Case
	upCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//Numbers
	numbers = "0123456789"
	//Special Characters
	specialChar = "@#$&*"
)

type CreatePasswordDTO struct {
	LowCaseQuantity     uint
	UpCaseQuantity      uint
	NumbersQuantity     uint
	SpecialCharQuantity uint
}

// Function to generate a random password
func Generate(p *CreatePasswordDTO) (password string) {
	var str = ""
	for i := uint(0); i < p.LowCaseQuantity; i++ {
		randNum := rand.Intn(len(lowCase))
		str += string(lowCase[randNum])
	}

	for i := uint(0); i < p.UpCaseQuantity; i++ {
		randNum := rand.Intn(len(upCase))
		str += string(upCase[randNum])
	}

	for i := uint(0); i < p.NumbersQuantity; i++ {
		randNum := rand.Intn(len(numbers))
		str += string(numbers[randNum])
	}

	for i := uint(0); i < p.SpecialCharQuantity; i++ {
		randNum := rand.Intn(len(specialChar))
		str += string(specialChar[randNum])
	}

	inRune := []rune(str)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	password = string(inRune)

	return
}
