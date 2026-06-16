package service

import (
	"fmt"
	"os"
	"password-manager/internal/crypto"
	"password-manager/internal/models"
	"strings"

	"golang.org/x/term"
)

func (s *service) Create() {
	var user models.Service

	fmt.Print("Введите название сервиса: ")
	serviceName, err := s.reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения названия сервиса:", err)
		return
	}
	user.ServiceName = strings.TrimSpace(serviceName)
	if user.ServiceName == "" {
		fmt.Println("Название сервиса не может быть пустым")
		return
	}

	fmt.Print("Введите логин пользователя: ")
	username, err := s.reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения логина:", err)
		return
	}
	user.User.Username = strings.TrimSpace(username)
	if user.User.Username == "" {
		fmt.Println("Логин не может быть пустым")
		return
	}

	fmt.Print("Введите пароль: ")
	plainPassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		fmt.Println("Ошибка чтения пароля:", err)
		return
	}

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
}
