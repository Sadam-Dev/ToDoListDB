package db

import "errors"

func Migrate() error {
	err := ConnectToDB()
	if err != nil {
		return err
	}
	tx, err := dbConn.Begin()
	if err != nil {
		return errors.New("Unable to begin transaction: " + err.Error())
	}

	_, err = tx.Exec(CreateTasksTableQuery)
	if err != nil {
		tx.Rollback()
		return errors.New("Unable to create tasks table: " + err.Error())
	}

	err = tx.Commit()

	if err != nil {
		return errors.New("Unable to commit transaction: " + err.Error())
	}
	return nil

}
