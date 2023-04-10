package main

import (
	"fmt"
	"log"
	"net/http"
	"social-forum/api/controllers"
	"social-forum/api/repository"
	"social-forum/api/services"
)

func main() {
	db := repository.LoadConfig()
	db = repository.NewCreateTables(db)

	repos := repository.NewRepository(*db)
	services := services.NewService(repos)
	handlers := controllers.NewHandler(services)

	fmt.Println(services)
	log.Println("Запуск веб-сервера на http://localhost:8080/ ")
	err := http.ListenAndServe(":8080", handlers.InitRoutes())

	if err != nil {
		fmt.Println(http.StatusServiceUnavailable)
		return
	}
}
