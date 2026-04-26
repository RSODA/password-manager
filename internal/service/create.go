package service

import (
	"fmt"
	"os"
	"password-manager/internal/crypto"
	"password-manager/internal/models"

	"golang.org/x/term"
)

func (s *service) Create() {
	var user models.Service

	fmt.Println("Введите название сервиса ")
	fmt.Scan(&user.ServiceName)
	fmt.Println("Введите login пользователя")
	fmt.Scan(&user.User.Username)
	fmt.Println("Введите пароль")
	plainPassword, err := term.ReadPassword(int(os.Stdin.Fd()))

	result, err := crypto.Encrypt(plainPassword, s.mk)
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

	s.DefaultWindow()
}
