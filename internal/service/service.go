package service

import "password-manager/internal/repo"

type Service interface {
	Create(mk []byte)
	Get(mk []byte)
}

type service struct {
	Repo repo.Repo
}

func New(r repo.Repo) Service {
	return &service{
		Repo: r,
	}
}
