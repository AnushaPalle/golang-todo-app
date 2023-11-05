package main

import (
    "github.com/gin-gonic/gin"
    "github.com/AnushaPalle/golang-todo-app/routes"
)

func main() {
    r := gin.Default()
    routes.SetupRoutes(r)
    r.Run(":8080")
}


// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// )

// var db *gorm.DB
// var err error

// type Todo struct {
// 	ID        uint   `json:"id" gorm:"primary_key"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

// func main() {
// 	// Connect to the PostgreSQL database
// 	db, err = gorm.Open("postgres", "host=localhost user=root password=secret dbname=todo_db sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Automigrate the Todo model
// 	db.AutoMigrate(&Todo{})

// 	// Initialize Gin router
// 	r := gin.Default()

// 	// Define routes
// 	r.GET("/todos", GetTodos)
// 	r.GET("/todos/:id", GetTodo)
// 	r.POST("/todos", CreateTodo)
// 	r.PUT("/todos/:id", UpdateTodo)
// 	r.DELETE("/todos/:id", DeleteTodo)

// 	// Run the application on port 8080
// 	r.Run(":8080")
// }

// // GetTodos returns a list of all todos
// func GetTodos(c *gin.Context) {
// 	var todos []Todo
// 	db.Find(&todos)
// 	c.JSON(http.StatusOK, todos)
// }

// // GetTodo retrieves a specific todo by ID
// func GetTodo(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var todo Todo
// 	err := db.Where("id = ?", id).First(&todo).Error
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, todo)
// 	}
// }

// // CreateTodo creates a new todo
// func CreateTodo(c *gin.Context) {
// 	var todo Todo
// 	c.BindJSON(&todo)
// 	db.Create(&todo)
// 	c.JSON(http.StatusCreated, todo)
// }

// // UpdateTodo updates a todo by ID
// func UpdateTodo(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var todo Todo
// 	err := db.Where("id = ?", id).First(&todo).Error
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 		return
// 	}
// 	c.BindJSON(&todo)
// 	db.Save(&todo)
// 	c.JSON(http.StatusOK, todo)
// }

// // DeleteTodo deletes a todo by ID
// func DeleteTodo(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var todo Todo
// 	err := db.Where("id = ?", id).First(&todo).Error
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 		return
// 	}
// 	db.Delete(&todo)
// 	c.JSON(http.StatusOK, gin.H{"id #" + id: "deleted"})
// }
