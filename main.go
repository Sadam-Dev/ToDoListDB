package main

import (
	"fmt"
	"log"
	"todoList/controllers"
	"todoList/db"
)

func main() {
	if err := db.ConnectToDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		if err := db.CloseDbConnection(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	for {
		controllers.PrintMenu()
		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			controllers.AddTask()
		case "2":
			controllers.PrintALLTasks()
		case "3":
			controllers.EditTaskByID()
		case "4":
			controllers.MarkTaskCompleted()
		case "5":
			controllers.DelTaskByID()
		case "exit":
			fmt.Println("Ради были помочь =)")
			return
		default:
			fmt.Println("Вы ввели несуществующую команду.")
		}
	}
}
