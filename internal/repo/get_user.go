package repo

import (
	"encoding/json"
	"fmt"
	"os"
	"password-manager/internal/models"
)

func (r *repo) GetUser(service string) (*models.User, error) {
	var services map[string]models.User

	res, err := os.ReadFile(r.Filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &services)
	if err != nil {
		return nil, err
	}

	svc, ok := services[service]
	if !ok {
		return nil, fmt.Errorf("service not found")
	}

	return &models.User{
		Username: svc.Username,
		Password: svc.Password,
	}, nil
}
