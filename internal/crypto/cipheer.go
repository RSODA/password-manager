package crypto

import (
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/pedroalbanese/gogost/gost3412128"
)

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	if len(key) != gost3412128.KeySize {
		return nil, fmt.Errorf("invalid key size")
	}

	block := gost3412128.NewCipher(key)
	ciphertext := make([]byte, block.BlockSize()+len(plaintext))

	nonce := ciphertext[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(ciphertext[block.BlockSize():], plaintext)

	return ciphertext, nil
}

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	if len(key) != gost3412128.KeySize {
		return nil, fmt.Errorf("invalid key size")
	}

	block := gost3412128.NewCipher(key)
	if len(ciphertext) < block.BlockSize() {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce := ciphertext[:block.BlockSize()]
	ciphertext = ciphertext[block.BlockSize():]

	plaintext := make([]byte, len(ciphertext))

	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
