package main

import (
    "github.com/amrchnk/todo-app"
    "github.com/amrchnk/todo-app/pkg/handler"
    "log"
)

func main(){
    handlers:=new(handler.Handler)

    srv:=new(todo.Server)
    if err:=srv.Run("8000",handlers.InitRoutes());err!=nil{
        log.Fatalf("Error with runing http-server: %s",err.Error())
    }
}