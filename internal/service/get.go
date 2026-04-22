package service

import (
	"fmt"
	"password-manager/internal/crypto"
	"password-manager/internal/models"
)

func (s *service) Get(mk []byte) {
	var servicename string
	var response models.Response

	fmt.Println("Напишите название сервиса")
	fmt.Scan(&servicename)

	res, err := s.Repo.GetUser(servicename)
	if err != nil {
		fmt.Println(err)
		return
	}

	decrypted, err := crypto.Decrypt(res.Password, mk)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(decrypted))

	fmt.Println()

	response.Username = res.Username
	response.Password = string(decrypted)

	fmt.Println(response)

}
