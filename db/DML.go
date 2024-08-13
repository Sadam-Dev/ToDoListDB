package db

const (
	CreatingTask = `INSERT INTO tasks (title, description)
					VALUES ($1, $2);`
	GetAllTasks = `SELECT tasks.id, 
       						tasks.title, 
       						tasks.description, 
       						tasks.status, 
       						tasks.created_at 
					FROM tasks
					WHERE tasks.is_deleted = FALSE;`
	GetTaskByID = `SELECT id, 
       						title, 
       						description, 
       						status, 
       						created_at 
					FROM tasks 
					WHERE id = $1;`
	UpdateTaskByID = `UPDATE tasks
						SET title = $1,
    						description = $2,
    						status = $3
						WHERE id = $4;`
	MarkTaskCompleted = `UPDATE tasks 
						SET status = TRUE
						WHERE id = $1;`

	SoftDeleteTaskByID = `UPDATE tasks SET is_deleted = true WHERE id = $1;`
)
