package service

import (
	"fmt"
	"password-manager/internal/clipboard"
	"password-manager/internal/crypto"
	"password-manager/internal/models"
	"strings"
	"time"
)

func (s *service) Get() {
	var response models.Response

	fmt.Print("Введите название сервиса: ")
	serviceName, err := s.reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения названия сервиса:", err)
		return
	}
	serviceName = strings.TrimSpace(serviceName)
	if serviceName == "" {
		fmt.Println("Название сервиса не может быть пустым")
		return
	}

	res, err := s.Repo.GetUser(serviceName)
	if err != nil {
		fmt.Println(err)
		return
	}

	decrypted, err := crypto.Decrypt(res.Password, s.mk)
	if err != nil {
		fmt.Println(err)
		return
	}

	response.Username = res.Username
	response.Password = string(decrypted)

	fmt.Println("Логин:", response.Username)

	if err := clipboard.CopyText(response.Password); err != nil {
		fmt.Println("Буфер обмена недоступен:", err)
		fmt.Println("Пароль:", response.Password)
		return
	}

	copiedAt := time.Now()
	clearAt := copiedAt.Add(ClipboardTimeout)

	fmt.Printf("Пароль скопирован в буфер обмена в %s и будет очищен через %d секунд (%s)\n",
		copiedAt.Format("15:04:05"),
		int(ClipboardTimeout.Seconds()),
		clearAt.Format("15:04:05"),
	)

	time.AfterFunc(ClipboardTimeout, func() {
		_ = clipboard.Clear()
	})
}
