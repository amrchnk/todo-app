package main

import (
    _ "github.com/lib/pq"
    "github.com/amrchnk/todo-app"
    "github.com/amrchnk/todo-app/pkg/handler"
    "github.com/amrchnk/todo-app/pkg/repository"
    "github.com/amrchnk/todo-app/pkg/service"
    "github.com/spf13/viper"
    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"
    "os"
)

func main(){
    logrus.SetFormatter(new(logrus.JSONFormatter))
    if err:=initConfig(); err!=nil{
        logrus.Fatalf("error initializing config: %s",err.Error())
    }

    if err:=godotenv.Load(); err!=nil{
        logrus.Fatalf("error loading env vars: %s",err.Error())
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
        logrus.Fatalf("failed to initialized db: %s",err.Error())
    }

    repos:=repository.NewRepository(db)
    services:=service.NewService(repos)
    handlers:=handler.NewHandler(services)

    srv:=new(todo.Server)
    if err:=srv.Run(viper.GetString("port"),handlers.InitRoutes());err!=nil{
        logrus.Fatalf("Error with runing http-server: %s",err.Error())
    }
}

func initConfig() error{
    viper.AddConfigPath("configs")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}