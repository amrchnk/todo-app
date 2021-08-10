package main

import (
    "github.com/amrchnk/todo-app"
    "github.com/amrchnk/todo-app/pkg/handler"
    "github.com/amrchnk/todo-app/pkg/repository"
    "github.com/amrchnk/todo-app/pkg/service"
    "github.com/spf13/viper"
    "log"
)

func main(){
    if err:=initConfig(); err!=nil{
        log.Fatalf("error initializing config: %s",err.Error())
    }
    repos:=repository.NewRepository()
    services:=service.NewService(repos)
    handlers:=handler.NewHandler(services)

    srv:=new(todo.Server)
    if err:=srv.Run(viper.GetString("8000"),handlers.InitRoutes());err!=nil{
        log.Fatalf("Error with runing http-server: %s",err.Error())
    }
}

func initConfig() error{
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}