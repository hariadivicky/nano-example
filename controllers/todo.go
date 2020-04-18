package controllers

import (
	"net/http"
	"strconv"

	"todo/models"

	"github.com/hariadivicky/nano"
)

// Todo Controller struct.
type Todo struct{}

// Index is functions to handle todo list request.
// route: GET /todos
func (ctrl *Todo) Index(c *nano.Context) {
	todo := new(models.Todo)

	collection := todo.All()

	c.JSON(http.StatusOK, nano.H{
		"collection": collection,
	})
}

// Store new todo to database.
// route: POST /todos
func (ctrl *Todo) Store(c *nano.Context) {
	title := c.PostForm("title")
	// accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False. Any other value are marked as false.
	isDone, _ := strconv.ParseBool(c.PostForm("is_done"))

	todo := (&models.Todo{
		Title:  title,
		IsDone: isDone,
	}).Save()

	c.JSON(http.StatusOK, nano.H{
		"todo": todo,
	})
}

// findTodoByParamId is functions to get todo by given parameter id.
func findTodoByParamID(c *nano.Context) *models.Todo {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 8)
	model := new(models.Todo)

	return model.Find(uint(id))
}

// Show todo by given parameter id.
// route: GET /todos/:id.
func (ctrl *Todo) Show(c *nano.Context) {
	todo := findTodoByParamID(c)

	// send http not found when todo does not exists.
	if todo == nil {
		c.String(http.StatusNotFound, "todo not found")
		return
	}

	c.JSON(http.StatusOK, nano.H{
		"todo": todo,
	})
}

// Toggle is functions to invers current done status at todo.
// route: PUT /todos/:id/toggle
func (ctrl *Todo) Toggle(c *nano.Context) {
	todo := findTodoByParamID(c)

	// send http not found when todo does not exists.
	if todo == nil {
		c.String(http.StatusNotFound, "todo not found")
		return
	}

	updated := todo.ToggleDoneStatus()

	c.JSON(http.StatusOK, nano.H{
		"updated": updated,
		"todo":    todo,
	})
}

// Destroy is functions to delete todo from database.
// route: DELETE /todos/:id
func (ctrl *Todo) Destroy(c *nano.Context) {
	todo := findTodoByParamID(c)

	// send http not found when todo does not exists.
	if todo == nil {
		c.String(http.StatusNotFound, "todo not found")
		return
	}

	deleted := todo.Delete()

	c.JSON(http.StatusOK, nano.H{
		"deleted": deleted,
	})
}
