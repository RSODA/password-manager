package service

import (
	"fmt"
	"password-manager/internal/crypto"
	"password-manager/internal/models"
)

func (s *service) Create(mk []byte) {
	var user models.Service
	var plainPassword string

	fmt.Println("Введите название сервиса ")
	fmt.Scan(&user.ServiceName)
	fmt.Println("Введите login пользователя")
	fmt.Scan(&user.User.Username)
	fmt.Println("Введите пароль")
	fmt.Scan(&plainPassword)

	result, err := crypto.Encrypt([]byte(plainPassword), mk)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := s.Repo.SaveUser(&models.Service{
		ServiceName: user.ServiceName,
		User: models.User{
			Username: user.User.Username,
			Password: result,
		},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
