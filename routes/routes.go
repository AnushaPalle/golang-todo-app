package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/AnushaPalle/golang-todo-app/controller"
    "github.com/AnushaPalle/golang-todo-app/service"
    "github.com/AnushaPalle/golang-todo-app/repository"
    "github.com/AnushaPalle/golang-todo-app/db"
)

func SetupRoutes(r *gin.Engine) {
    db, err := database.Init()
    if err != nil {
        panic("Failed to connect to database")
    }

    tr := repository.NewTodoRepository(db)
    ts := service.NewTodoService(tr)
    tc := controller.NewTodoController(ts)

    v1 := r.Group("/api/v1/todos")
    {
        v1.POST("/", tc.CreateTodo)
        v1.GET("/", tc.GetTodos)
        
	    // v1.GET("/:id", tc.GetTodo)
	    // v1.PUT("/:id", tc.UpdateTodo)
	    // v1.DELETE("/:id", tc.DeleteTodo)


        v1.GET("/:id", tc.GetTodoByID)
		v1.PUT("/:id", tc.UpdateTodoByID)
		v1.DELETE("/:id", tc.DeleteTodoByID)

        // v1.POST("/", tc.CreateTodos)
    }
}
