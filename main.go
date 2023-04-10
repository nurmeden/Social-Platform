package main

import (
	"fmt"
	"log"
	"net/http"
	"social-forum/internal/controllers"
	"social-forum/internal/repository"
	"social-forum/internal/usecase"
)

func main() {
	db := repository.LoadConfig()
	db = repository.NewCreateTables(db)

	repos := repository.NewRepository(*db)
	services := usecase.NewUseCase(repos)
	fmt.Println("ok")
	handlers := controllers.NewHandler(services)

	fmt.Println(services)
	log.Println("Запуск веб-сервера на http://localhost:8080/ ")
	err := http.ListenAndServe(":8080", handlers.InitRoutes())

	if err != nil {
		fmt.Println(http.StatusServiceUnavailable)
		return
	}
}
