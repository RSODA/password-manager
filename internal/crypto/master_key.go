package crypto

import (
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func getOrCreateSalt(path string, saltLen int) ([]byte, error) {
	salt, err := os.ReadFile(path)
	if err == nil {
		if len(salt) != saltLen {
			return nil, errors.New("invalid salt length")
		}
		return salt, nil
	}

	if !os.IsNotExist(err) {
		return nil, err
	}

	salt = make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	if err := os.WriteFile(path, salt, 0600); err != nil {
		return nil, err
	}

	return salt, nil
}

func GetMasterKey(password []byte, saltPath string) ([]byte, error) {

	iteration := 15000 + (9 * 5000)
	saltLen := 16 + (2006 % 8)
	salt, err := getOrCreateSalt(saltPath, saltLen)
	if err != nil {
		return nil, err
	}

	mk := pbkdf2.Key([]byte(password), salt, iteration, 32, sha512.New)

	return mk, nil
}
