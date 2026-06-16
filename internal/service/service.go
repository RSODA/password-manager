package service

import (
	"bufio"
	"io"
	"os"
	"password-manager/internal/repo"
	"time"
)

const ClipboardTimeout = 20 * time.Second

type Service interface {
	Create()
	Get()
	DefaultWindow()
}

type service struct {
	Repo   repo.Repo
	mk     []byte
	reader *bufio.Reader
}

func New(r repo.Repo, mk []byte) Service {
	return NewWithReader(r, mk, os.Stdin)
}

func NewWithReader(r repo.Repo, mk []byte, input io.Reader) Service {
	return &service{
		Repo:   r,
		mk:     mk,
		reader: bufio.NewReader(input),
	}
}
