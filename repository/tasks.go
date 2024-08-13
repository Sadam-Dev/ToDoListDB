package repository

import (
	"todoList/db"
	"todoList/models"
)

func AddTask(t models.Task) error {
	_, err := db.GetDBConn().Exec(db.CreatingTask, t.Title, t.Description)
	if err != nil {
		return err
	}
	return nil
}

func GetTasks() (tasks []models.Task, err error) {
	rows, err := db.GetDBConn().Query(db.GetAllTasks)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var t models.Task
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, err
}

func GetTaskByID(id int) (t models.Task, err error) {
	row := db.GetDBConn().QueryRow(db.GetTaskByID, id)

	err = row.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt)
	if err != nil {
		return models.Task{}, err
	}
	return t, nil
}

func UpdateTask(t models.Task) error {
	_, err := db.GetDBConn().Exec(db.UpdateTaskByID, t.Title, t.Description, t.Status, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func CompleteTask(id int) error {
	_, err := db.GetDBConn().Exec(db.MarkTaskCompleted, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(id int) error {
	_, err := db.GetDBConn().Exec(db.SoftDeleteTaskByID, id)
	if err != nil {
		return err
	}
	return nil
}
