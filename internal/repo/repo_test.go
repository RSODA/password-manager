package repo

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"password-manager/internal/models"
)

func TestSaveAndGetUser(t *testing.T) {
	path := filepath.Join(t.TempDir(), "vault.json")
	repository := NewRepo(path)
	password := []byte{1, 2, 3, 4}

	_, err := repository.SaveUser(&models.Service{
		ServiceName: "mail",
		User: models.User{
			Username: "user@example.com",
			Password: password,
		},
	})
	if err != nil {
		t.Fatalf("SaveUser() error = %v", err)
	}

	got, err := repository.GetUser("mail")
	if err != nil {
		t.Fatalf("GetUser() error = %v", err)
	}
	if got.Username != "user@example.com" {
		t.Fatalf("Username = %q, want user@example.com", got.Username)
	}
	if !bytes.Equal(got.Password, password) {
		t.Fatalf("Password = %v, want %v", got.Password, password)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}
	if info.Mode().Perm() != 0600 {
		t.Fatalf("file permissions = %v, want 0600", info.Mode().Perm())
	}
}

func TestGetUserNotFound(t *testing.T) {
	path := filepath.Join(t.TempDir(), "vault.json")
	repository := NewRepo(path)

	_, err := repository.SaveUser(&models.Service{
		ServiceName: "mail",
		User: models.User{
			Username: "user@example.com",
			Password: []byte{1, 2, 3, 4},
		},
	})
	if err != nil {
		t.Fatalf("SaveUser() error = %v", err)
	}

	if _, err := repository.GetUser("unknown"); err == nil {
		t.Fatal("GetUser() error = nil, want service not found")
	}
}
