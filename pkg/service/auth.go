package service

import (
    "crypto/sha1"
    "github.com/amrchnk/todo-app/pkg/repository"
    "github.com/amrchnk/todo-app"
    "fmt"
)

const salt="jorhpihipbfqiepbgiergbqpier"

type AuthService struct{
    repo repository.Authorization
}

func NewAuthService(repo repository.Authorization)*AuthService{
    return &AuthService{repo:repo}
}

func (s *AuthService) CreateUser(user todo.User)(int,error){
    user.Password=generatePasswordHash(user.Password)
    return s.repo.CreateUser(user)
}

func generatePasswordHash(password string)string{
    hash:=sha1.New()
    hash.Write([]byte(password))

    return fmt.Sprintf("%x",hash.Sum([]byte(salt)))
}