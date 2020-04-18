package models

import (
	"database/sql"
)

// Todo model.
type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"is_done"`
}

// All functions to returns all todo list.
func (model *Todo) All() []Todo {
	rows, err := db.Query("SELECT * FROM todos")
	handleError("query failed: %v", err)

	var todos []Todo
	defer rows.Close()

	for rows.Next() {
		var todo Todo

		err := rows.Scan(&todo.ID, &todo.Title, &todo.IsDone)
		handleError("scan failed: %v", err)

		todos = append(todos, todo)
	}

	return todos
}

// Save todo.
func (model *Todo) Save() *Todo {
	stmt, err := db.Prepare("INSERT INTO todos (title, is_done) VALUES (?, ?)")
	handleError("could prepare statement: %v", err)

	res, err := stmt.Exec(model.Title, model.IsDone)
	handleError("failed to store: %v", err)

	id, _ := res.LastInsertId()
	model.ID = uint(id)

	return model
}

// Find todo by id.
func (model *Todo) Find(id uint) *Todo {
	row := db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	todo := &Todo{}
	err := row.Scan(&todo.ID, &todo.Title, &todo.IsDone)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
	}

	return todo
}

// ToggleDoneStatus is functions to invers current is_done value.
func (model *Todo) ToggleDoneStatus() bool {
	model.IsDone = !model.IsDone

	stmt, err := db.Prepare("UPDATE todos SET is_done = ? WHERE id = ?")
	handleError("could not prepare statement: %v", err)

	res, err := stmt.Exec(model.IsDone, model.ID)
	handleError("query failed: %v", err)

	affecteds, err := res.RowsAffected()
	handleError("could not get affected rows: %v", err)

	if affecteds > 0 {
		return true
	}

	return false
}

// Delete is functions to remove todo from database.
func (model *Todo) Delete() bool {
	stmt, err := db.Prepare("DELETE FROM todos WHERE id = ?")
	handleError("could not prepare statement: %v", err)

	res, err := stmt.Exec(model.ID)
	handleError("query failed: %v", err)

	affecteds, err := res.RowsAffected()
	handleError("could not get affected rows: %v", err)

	if affecteds > 0 {
		return true
	}

	return false
}
