package service

import (
	"inventory/model"
	"inventory/repository"
)

type TodoService interface {
	CreateTodo(todo *model.Todo) error
	GetAllTodos() ([]model.Todo, error)
	GetTodoByID(id string) (model.Todo, error)
	UpdateTodoByID(todo *model.Todo) error
	DeleteTodoByID(id string) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

func (s *todoService) CreateTodo(todo *model.Todo) error {
	return s.repo.Create(todo)
}

func (s *todoService) GetAllTodos() ([]model.Todo, error) {
	return s.repo.FindAll()
}

func (s *todoService) GetTodoByID(id string) (model.Todo, error) {
	return s.repo.FindByID(id)
}

func (s *todoService) UpdateTodoByID(todo *model.Todo) error {
	return s.repo.Update(todo)
}

func (s *todoService) DeleteTodoByID(id string) error {
	return s.repo.DeleteByID(id)
}
