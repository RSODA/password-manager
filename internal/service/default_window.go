package service

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *service) DefaultWindow() {
	for {
		fmt.Print("\nВыберите действие:\n" +
			"1. Добавить пароль\n" +
			"2. Найти пароль\n" +
			"3. Выход\n" +
			"> ")

		input, err := s.reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения команды:", err)
			return
		}

		count, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Ошибка, введите номер действия")
			continue
		}

		switch count {
		case 1:
			s.Create()
		case 2:
			s.Get()
		case 3:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Ошибка, введите правильное значение")
		}
	}
}
