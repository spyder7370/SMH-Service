package encrypt

import "golang.org/x/crypto/bcrypt"

func Hash(input string) (string, error) {
	hashedInput, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedInput), nil
}

func Verify(hashedInput, input string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedInput), []byte(input))
	return err != nil
}
