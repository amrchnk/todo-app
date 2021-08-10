package main

import (
    _ "github.com/lib/pq"
    "github.com/amrchnk/todo-app"
    "github.com/amrchnk/todo-app/pkg/handler"
    "github.com/amrchnk/todo-app/pkg/repository"
    "github.com/amrchnk/todo-app/pkg/service"
    "github.com/spf13/viper"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func main(){
    if err:=initConfig(); err!=nil{
        log.Fatalf("error initializing config: %s",err.Error())
    }

    if err:=godotenv.Load(); err!=nil{
            log.Fatalf("error loading env vars: %s",err.Error())
    }

    db,err:=repository.NewPostgresDB(repository.Config{
        Host: viper.GetString("db.host"),
        Port: viper.GetString("db.port"),
        Username: viper.GetString("db.username"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName: viper.GetString("db.dbname"),
        SSLMode: viper.GetString("db.sslmode"),
    })
    if err!=nil{
        log.Fatalf("failed to initialized db: %s",err.Error())
    }

    repos:=repository.NewRepository(db)
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