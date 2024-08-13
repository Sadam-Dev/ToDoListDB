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

func AddTask() error {
	var task models.Task

	fmt.Println("Введите заголовок задачи:")
	reader := bufio.NewReader(os.Stdin)
	task.Title, _ = reader.ReadString('\n')
	task.Title = strings.TrimSpace(task.Title)

	fmt.Println("Введите описание задачи:")
	reader = bufio.NewReader(os.Stdin)
	task.Description, _ = reader.ReadString('\n')
	task.Description = strings.TrimSpace(task.Description)

	err := repository.AddTask(task)
	if err != nil {
		return err
	}
	fmt.Println("Задача успешно добавлена")
	return nil
}

func PrintALLTasks() {
	fmt.Printf("ID | Title | Description | Status | CreatedAt\n")
	fmt.Println("---------------------------------------------------------")
	tasks, err := repository.GetTasks()
	if err != nil {
		fmt.Println("Error during getting list of operations:", err.Error())
	}

	for _, t := range tasks {
		fmt.Printf("%d | %s | %s | %t | %s\n",
			t.ID,
			t.Title,
			t.Description,
			t.Status,
			t.CreatedAt.Format(time.RFC3339))

	}
}

func EditTaskByID() {
	var id int
	fmt.Println("Введите ID вашей задачи")
	fmt.Scan(&id)

	task, err := repository.GetTaskByID(id)
	if err != nil {
		fmt.Println("Ошибка при получении задачи:", err)
		return
	}

	fmt.Printf("Текущий заголовок задачи: %s\n", task.Title)
	fmt.Println("Введите новый заголовок задачи:")
	reader := bufio.NewReader(os.Stdin)
	newTitle, _ := reader.ReadString('\n')
	task.Title = strings.TrimSpace(newTitle)

	fmt.Printf("Текущее описание задачи: %s\n", task.Description)
	fmt.Println("Введите новое описание задачи:")
	newDescription, _ := reader.ReadString('\n')
	task.Description = strings.TrimSpace(newDescription)

	fmt.Printf("Текущий статус задачи: %t\n", task.Status)
	fmt.Println("Введите новый статус задачи (true для выполнено, false для не выполнено):")
	var newStatus bool
	fmt.Scanln(&newStatus)
	task.Status = newStatus

	err = repository.UpdateTask(task)
	if err != nil {
		fmt.Println("Ошибка при обновлении задачи:", err)
		return
	}

	fmt.Println("Задача успешно обновлена")

}

func MarkTaskCompleted() error {
	fmt.Println("Введите ID задачи для отметки как выполненной:")
	var id int
	fmt.Scanln(&id)

	err := repository.CompleteTask(id)
	if err != nil {
		return fmt.Errorf("ошибка при отметке задачи как выполненной: %v", err)
	}

	fmt.Println("Задача успешно отмечена как выполненная")
	return nil
}

func DelTaskByID() {
	var id int
	fmt.Printf("Введите ID задачи которую хотите удалить.\nID: ")
	_, _ = fmt.Scan(&id)
	err := repository.DeleteTask(id)
	if err != nil {
		fmt.Println("Error during deleting operation", err.Error())
	}
	fmt.Println("Задача успешно удалена")
}
