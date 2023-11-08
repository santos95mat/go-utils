package passwordGenerator

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

// Function to generate a random password
func Generate(passwordLen int) string {
	password := ""

	for i := 0; i < passwordLen; i++ {
		randNum := rand.Intn(4)

		switch randNum {
		case 0:
			randNum := rand.Intn(len(lowCase))
			password += string(lowCase[randNum])
			break
		case 1:
			randNum := rand.Intn(len(upCase))
			password += string(upCase[randNum])
			break
		case 2:
			randNum := rand.Intn(len(numbers))
			password += string(numbers[randNum])
			break
		case 3:
			randNum := rand.Intn(len(specialChar))
			password += string(specialChar[randNum])
			break
		}

	}

	return password
}
