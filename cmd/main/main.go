package main

import (
	"fmt"
	"password-manager/internal/crypto"
	repo2 "password-manager/internal/repo"
	"password-manager/internal/service"
)

func main() {
	var input string
	var count int64

	min_len := 8 + (6 % 5)

	fmt.Println("Введите мастер пароль: ")
	fmt.Scan(&input)

	if len(input) < min_len {
		fmt.Println("Ваша длина мастер пароля меньше: ", min_len)
	}

	mk, err := crypto.GetMasterKey(input, "master.salt")
	if err != nil {
		fmt.Println("Не удалось получить мастер-ключ:", err)
		return
	}
	if err := crypto.ValidateOrCreateKeyCheck("master.keycheck", mk); err != nil {
		fmt.Println("Мастер-пароль не совпадает с текущим хранилищем:", err)
		return
	}

	repo := repo2.NewRepo("vault_06092006.json")
	s := service.New(repo)

	fmt.Print("Выебрите действие: \n" +
		"1. Добавить пароль\n" +
		"2. Найти пароль")
	fmt.Scan(&count)

	if count == 1 {
		s.Create(mk)
	} else if count == 2 {
		s.Get(mk)
	}
}
