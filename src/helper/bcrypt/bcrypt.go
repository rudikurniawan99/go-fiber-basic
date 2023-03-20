package bcrypt_helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return hash, err
}

func ComparePassword(pass string, hashPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPassword, []byte(pass))
	return err
}
