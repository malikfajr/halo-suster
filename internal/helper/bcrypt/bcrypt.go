package bcrypt

import (
	"github.com/malikfajr/halo-suster/config"
	bc "golang.org/x/crypto/bcrypt"
)

func CreateHash(password string) string {
	p, err := bc.GenerateFromPassword([]byte(password), config.Bcrypt.Salt)
	if err != nil {
		panic(err)
	}

	return string(p)
}

func PasswordIsValid(password string, hash string) bool {
	err := bc.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
