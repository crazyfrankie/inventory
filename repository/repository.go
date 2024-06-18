package repository

import (
	"inventory/config"
	"inventory/model"
)

type TodoRepository interface {
	Create(todo *model.Todo) error
	FindAll() ([]model.Todo, error)
	FindByID(id string) (model.Todo, error)
	Update(todo *model.Todo) error
	DeleteByID(id string) error
}

type todoRepository struct{}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (r *todoRepository) Create(todo *model.Todo) error {
	return config.DB.Create(todo).Error
}

func (r *todoRepository) FindAll() ([]model.Todo, error) {
	var todoList []model.Todo
	err := config.DB.Find(&todoList).Error
	return todoList, err
}

func (r *todoRepository) FindByID(id string) (model.Todo, error) {
	var todo model.Todo
	err := config.DB.Where("id = ?", id).First(&todo).Error
	return todo, err
}

func (r *todoRepository) Update(todo *model.Todo) error {
	return config.DB.Save(todo).Error
}

func (r *todoRepository) DeleteByID(id string) error {
	return config.DB.Where("id = ?", id).Delete(&model.Todo{}).Error
}
