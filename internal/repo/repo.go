package repo

import "password-manager/internal/models"

type Repo interface {
	SaveUser(user *models.Service) (string, error)
	GetUser(service string) (*models.User, error)
}

type repo struct {
	Filename string
}

func NewRepo(filename string) Repo {
	return &repo{
		Filename: filename,
	}
}
