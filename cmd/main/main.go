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

const (
	accessStartHour = 9
	accessEndHour   = 18
	minPasswordLen  = 8 + (6 % 5)
)

func isAccessAllowed(now time.Time) bool {
	hour := now.Hour()
	return hour >= accessStartHour && hour < accessEndHour
}

func main() {
	now := time.Now()
	if !isAccessAllowed(now) {
		fmt.Printf("Доступ разрешён только с %02d:00 до %02d:00. Текущее время: %s\n",
			accessStartHour,
			accessEndHour,
			now.Format("15:04:05"),
		)
		os.Exit(1)
	}

	fmt.Print("Введите мастер-пароль: ")
	input, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		fmt.Println("Ошибка чтения мастер-пароля:", err)
		return
	}

	if len(input) < minPasswordLen {
		fmt.Println("Длина мастер-пароля меньше:", minPasswordLen)
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
