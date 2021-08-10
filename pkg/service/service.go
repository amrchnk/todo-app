package service

import (
    "github.com/amrchnk/todo-app/pkg/repository"
    "github.com/amrchnk/todo-app"
)

type Authorization interface{
    CreateUser(user todo.User)(int,error)
}

type TodoList interface{

}

type TodoItem interface{

}

type Service struct{
    Authorization
    TodoList
    TodoItem
}

func NewService(repos *repository.Repository) *Service{
    return &Service{
        Authorization: NewAuthService(repos.Authorization),
    }
}