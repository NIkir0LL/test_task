package main

import (
    "database/sql"
    "log"
    "user-api/internal/handler"
    "user-api/internal/repository"

    _ "github.com/lib/pq"
    "github.com/gin-gonic/gin"
)

func main() {
    // Подключение к PostgreSQL
    db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/userdb?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Инициализация репозитория и хендлеров
    repo := repository.NewUserRepository(db)
    userHandler := handler.NewUserHandler(repo)

    // Настройка роутера Gin
    router := gin.Default()
    
    router.POST("/users", userHandler.CreateUser)
    router.GET("/users/:id", userHandler.GetUser)
    router.PUT("/users/:id", userHandler.UpdateUser)

    // Запуск сервера
    if err := router.Run(":8080"); err != nil {
        log.Fatal(err)
    }
}