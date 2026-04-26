package service

import "fmt"

func (s *service) DefaultWindow() {
	var count int

	fmt.Print("Выебрите действие: \n" +
		"1. Добавить пароль\n" +
		"2. Найти пароль")
	fmt.Scan(&count)

	if count == 1 {
		s.Create()
	} else if count == 2 {
		s.Get()
	} else {
		fmt.Println("Ошибка, введите правильное значение!")
	}
}
