package repo

import (
	"encoding/json"
	"os"
	"password-manager/internal/models"
)

func (r *repo) SaveUser(user *models.Service) (string, error) {
	data := make(map[string]models.User)

	file, err := os.ReadFile(r.Filename)
	if err == nil {
		if err := json.Unmarshal(file, &data); err != nil {
			return "", err
		}
	} else if !os.IsNotExist(err) {
		return "", err
	}

	data[user.ServiceName] = models.User{
		Username: user.User.Username,
		Password: user.User.Password,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(r.Filename, jsonData, 0600); err != nil {
		return "", err
	}

	return user.ServiceName, nil
}
