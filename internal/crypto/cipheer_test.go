package crypto

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := bytes.Repeat([]byte{1}, 32)
	plaintext := []byte("secret-password")

	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Encrypt() error = %v", err)
	}
	if bytes.Contains(ciphertext, plaintext) {
		t.Fatal("ciphertext contains plaintext")
	}

	decrypted, err := Decrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("Decrypt() error = %v", err)
	}
	if !bytes.Equal(decrypted, plaintext) {
		t.Fatalf("Decrypt() = %q, want %q", decrypted, plaintext)
	}
}

func TestDecryptWithWrongKeyDoesNotReturnPlaintext(t *testing.T) {
	key := bytes.Repeat([]byte{1}, 32)
	wrongKey := bytes.Repeat([]byte{2}, 32)
	plaintext := []byte("secret-password")

	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Encrypt() error = %v", err)
	}

	decrypted, err := Decrypt(ciphertext, wrongKey)
	if err != nil {
		t.Fatalf("Decrypt() error = %v", err)
	}
	if bytes.Equal(decrypted, plaintext) {
		t.Fatal("Decrypt() with wrong key returned plaintext")
	}
}

func TestEncryptRejectsInvalidKey(t *testing.T) {
	if _, err := Encrypt([]byte("value"), []byte("short")); err == nil {
		t.Fatal("Encrypt() error = nil, want invalid key error")
	}
}
