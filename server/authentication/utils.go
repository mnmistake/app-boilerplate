package authentication

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]uint8, error) {
	bytePassword := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	
	return hashedPassword, nil
}