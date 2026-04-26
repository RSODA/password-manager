package main

import (
	"fmt"
	"os"
	"password-manager/internal/crypto"
	repo2 "password-manager/internal/repo"
	"password-manager/internal/service"
	"time"

	"golang.org/x/term"
)

func main() {
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)

	go func() {
		for {
			if time.Now().After(endTime) {
				fmt.Println("Конец")
				os.Exit(0)
			}
		}
	}()

	min_len := 8 + (6 % 5)

	fmt.Println("Введите мастер пароль: ")
	input, err := term.ReadPassword(int(os.Stdin.Fd()))
	if len(input) < min_len {
		fmt.Println("Ваша длина мастер пароля меньше: ", min_len)
		return
	}

	if len(input) < min_len {
		fmt.Println("Ваша длина мастер пароля меньше: ", min_len)
		return
	}

	mk, err := crypto.GetMasterKey(input, "master.salt")
	if err != nil {
		fmt.Println("Не удалось получить мастер-ключ:", err)
		return
	}

	repo := repo2.NewRepo("vault_06092006.json")
	s := service.New(repo, mk)

	s.DefaultWindow()
}
