package service

import (
    "github.com/AnushaPalle/golang-todo-app/models"
    "github.com/AnushaPalle/golang-todo-app/repository"
)

type TodoService struct {
    TodoRepository *repository.TodoRepository
}

func NewTodoService(tr *repository.TodoRepository) *TodoService {
    return &TodoService{TodoRepository: tr}
}

func (ts *TodoService) Create(todo *models.Todo) error {
    return ts.TodoRepository.Create(todo)
}

func (ts *TodoService) FindAll() ([]models.Todo, error) {
    return ts.TodoRepository.FindAll()
}

//hi

func (ts *TodoService) FindByID(id string) (*models.Todo, error) {
	return ts.TodoRepository.FindByID(id)
}

func (ts *TodoService) UpdateByID(id string, updatedTodo *models.Todo) error {
	return ts.TodoRepository.UpdateByID(id, updatedTodo)
}

func (ts *TodoService) DeleteByID(id string) error {
	return ts.TodoRepository.DeleteByID(id)
}

func (ts *TodoService) CreateMultiple(todos []models.Todo) error {
    for i := range todos {
        if err := ts.TodoRepository.Create(&todos[i]); err != nil {
            return err
        }
    }
    return nil
}
