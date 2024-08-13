package main

import (
	"fmt"
	"log"
	"todoList/controllers"
	"todoList/db"
)

func main() {
	err := db.Migrate()
	if err != nil {
		log.Fatal(err)
	}

	for {
		controllers.PrintMenu()
		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			err := controllers.AddTask()
			if err != nil {
				log.Fatal(err)
			}
		case "2":
			controllers.PrintALLTasks()
		case "3":
			controllers.EditTaskByID()
		case "4":
			controllers.MarkTaskCompleted()
		case "5":
			controllers.DelTaskByID()
		case "0":
			fmt.Println("Ради были помочь =)")
			return
		default:
			fmt.Println("Вы ввели несуществующую команду.")
		}
	}
}
