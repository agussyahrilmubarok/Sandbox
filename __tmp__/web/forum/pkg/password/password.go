package password

import "golang.org/x/crypto/bcrypt"

func Hash(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed)
}