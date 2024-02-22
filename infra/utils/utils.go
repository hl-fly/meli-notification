package utils

import (
	"crypto/rand"
	"math/big"

	"github.com/labstack/gommon/random"
	"golang.org/x/crypto/bcrypt"
)

func Hash(input string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func ValidateHash(plain string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(plain))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func SaltPassword(password string, salt string) string {
	if password == "" || salt == "" {
		return password
	}
	return password + "$" + salt
}

func RandomString(length uint) (string, error) {
	charset := []rune(random.Alphanumeric)

	result := make([]rune, length)
	maxIndex := big.NewInt(int64(len(charset)))
	for i := range result {
		n, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}
