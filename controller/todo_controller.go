package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/AnushaPalle/golang-todo-app/models"
    "github.com/AnushaPalle/golang-todo-app/service"
    "net/http"
)

type TodoController struct {
    TodoService *service.TodoService
}

func NewTodoController(ts *service.TodoService) *TodoController {
    return &TodoController{TodoService: ts}
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := tc.TodoService.Create(&todo); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, todo)
}

func (tc *TodoController) GetTodos(c *gin.Context) {
    todos, err := tc.TodoService.FindAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, todos)
}

// hi

func (tc *TodoController) GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := tc.TodoService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (tc *TodoController) UpdateTodoByID(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo models.Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.TodoService.UpdateByID(id, &updatedTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}

func (tc *TodoController) DeleteTodoByID(c *gin.Context) {
	id := c.Param("id")
	err := tc.TodoService.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (tc *TodoController) CreateTodos(c *gin.Context) {
    var todos []models.Todo

    if err := c.ShouldBindJSON(&todos); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := tc.TodoService.CreateMultiple(todos); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, todos)
}

