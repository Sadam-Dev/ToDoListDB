package repository

import (
	"todoList/db"
	"todoList/models"
)

func AddTask(task models.Task) error {
	return db.GetDBConn().Create(&task).Error
}

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := db.GetDBConn().Find(&tasks).Error
	return tasks, err
}

func UpdateTask(taskID int, title, description string) error {
	return db.GetDBConn().Model(&models.Task{}).
		Where("id = ?", taskID).Updates(models.Task{Title: title, Description: description}).Error
}

func CompleteTask(id int) error {
	var task models.Task
	err := db.GetDBConn().First(&task, id).Error
	if err != nil {
		return err
	}
	task.Status = !task.Status
	return db.GetDBConn().Save(&task).Error
}

func DeleteTask(id int) error {
	return db.GetDBConn().Delete(&models.Task{}, id).Error
}
