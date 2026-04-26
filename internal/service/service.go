package service

import "password-manager/internal/repo"

type Service interface {
	Create()
	Get()
	DefaultWindow()
}

type service struct {
	Repo repo.Repo
	mk   []byte
}

func New(r repo.Repo, mk []byte) Service {
	return &service{
		Repo: r,
		mk:   mk,
	}
}
