package repository

import (
    "github.com/AnushaPalle/golang-todo-app/models"
    "gorm.io/gorm"
)

type TodoRepository struct {
    DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
    return &TodoRepository{DB: db}
}

func (tr *TodoRepository) Create(todo *models.Todo) error {
    return tr.DB.Create(todo).Error
}

func (tr *TodoRepository) FindAll() ([]models.Todo, error) {
    var todos []models.Todo
    if err := tr.DB.Find(&todos).Error; err != nil {
        return nil, err
    }
    return todos, nil
}

// hi

func (tr *TodoRepository) FindByID(id string) (*models.Todo, error) {
	var todo models.Todo
	if err := tr.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (tr *TodoRepository) UpdateByID(id string, updatedTodo *models.Todo) error {
	return tr.DB.Model(&models.Todo{}).Where("id = ?", id).Updates(updatedTodo).Error
}

func (tr *TodoRepository) DeleteByID(id string) error {
	return tr.DB.Delete(&models.Todo{}, id).Error
}