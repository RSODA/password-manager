package service

import (
	"fmt"
	"password-manager/internal/crypto"
	"password-manager/internal/models"

	"github.com/atotto/clipboard"
)

func (s *service) Get() {
	var servicename string
	var response models.Response

	fmt.Println("Напишите название сервиса")
	fmt.Scan(&servicename)

	res, err := s.Repo.GetUser(servicename)
	if err != nil {
		fmt.Println(err)
		return
	}

	decrypted, err := crypto.Decrypt(res.Password, s.mk)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(decrypted))

	fmt.Println()

	response.Username = res.Username
	response.Password = string(decrypted)

	err = clipboard.WriteAll(response.Password)
	if err != nil {
		fmt.Println("ошибка при записи в буфер!", err)
		return
	}

	fmt.Println(response)

	s.DefaultWindow()
}
