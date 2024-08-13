package controllers

import "fmt"

func PrintMenu() {
	fmt.Println("Добро пожаловать в Todo-list")
	fmt.Println("Выберите команду: ")
	fmt.Println("1. Добавить задачу")
	fmt.Println("2. Получить список всех моих задач")
	fmt.Println("3. Редактирование задач")
	fmt.Println("4. Изменить статус задач")
	fmt.Println("5. Удалить задачу")
	fmt.Println("exit. Завершить программу")
}
