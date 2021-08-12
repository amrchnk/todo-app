package service

import (
    "github.com/amrchnk/todo-app/pkg/repository"
    "github.com/amrchnk/todo-app"
)

type TodoListService struct{
    repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList)*TodoListService{
    return &TodoListService{repo:repo}
}

func (s *TodoListService) Create(userId int,list todo.TodoList)(int,error){
    return s.repo.Create(userId,list)
}