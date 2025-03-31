package lib

import "golang.org/x/crypto/bcrypt"

// The purpose of this function is to encrypt password before store it in the database
func PasswordEncrypter(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return ""
	}
	return string(hashedPassword)
}

func PasswordDecrypter(userPassword, submittedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(submittedPassword)) == nil
}
