package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"todoList/models"
	"todoList/repository"
)

func AddTask() {
	var task models.Task

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите заголовок задачи: ")
	Title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении заголовка:", err)
		return
	}
	task.Title = strings.TrimSpace(Title)

	fmt.Println("Введите описание задачи:")

	Description, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении описания:", err)
		return
	}
	task.Description = strings.TrimSpace(Description)

	err = repository.AddTask(task)
	if err != nil {
		fmt.Println("Ошибка при добавлении задачи:", err)
		return
	}

	fmt.Println("Задача успешно добавлена")

}

func PrintALLTasks() {
	fmt.Printf("ID | Title | Description | CreatedAt\n")
	fmt.Println("----------------------------------------------------------------")
	tasks, err := repository.GetTasks()
	if err != nil {
		fmt.Println("Ошибка при получении задач:", err)
		return
	}

	for _, t := range tasks {

		fmt.Printf("%d | %s | %s | %s\n",
			t.ID,
			t.Title,
			t.Description,

			t.CreatedAt.Format(time.RFC3339))
		if t.Status == true {
			fmt.Println("Задача выполнена")
		} else {
			fmt.Println("Задача не выполнена")
		}

	}
}

func EditTaskByID() {
	var id int
	var newTitle, newDescription string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите ID задачи для редактирования:")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Ошибка при чтении ID задачи:", err)
		return
	}

	fmt.Println("Введите новый заголовок:")
	title, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении нового заголовка:", err)
		return
	}
	newTitle = strings.TrimSpace(title)

	fmt.Println("Введите новое описание:")
	description, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка при чтении нового описания:", err)
		return
	}
	newDescription = strings.TrimSpace(description)

	err = repository.UpdateTask(id, newTitle, newDescription)
	if err != nil {
		fmt.Println("Ошибка при обновлении задачи:", err)
		return
	}

	fmt.Println("Задача успешно обновлена!")

}

func MarkTaskCompleted() {
	var id int

	fmt.Println("Введите ID задачи для изменения статуса:")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Ошибка при чтении ID задачи:", err)
		return
	}

	err = repository.CompleteTask(id)
	if err != nil {
		fmt.Println("Ошибка при изменении статуса задачи:", err)
		return
	}

	fmt.Println("Статус задачи успешно изменен!")
}

func DelTaskByID() {
	var id int

	fmt.Printf("Введите ID задачи которую хотите удалить.\nID: ")
	_, err := fmt.Scan(&id)
	if err != nil {
		fmt.Println("Ошибка при чтении ID задачи: ")
		return
	}
	err = repository.DeleteTask(id)
	if err != nil {
		fmt.Println("Ошибка при удалении задачи: ", err)
	}
	fmt.Println("Задача успешно удалена")
}
