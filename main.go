package main

import (
	"fmt"
	"log"
	"net/http"
	"social-forum/api/repository"
	"social-forum/routes"
)

func main() {
	db := repository.LoadConfig()
	db = repository.NewCreateTables(db)
	log.Println("Запуск веб-сервера на http://localhost:8080/ ")
	err := http.ListenAndServe(":8080", routes.InitRoutes())

	if err != nil {
		fmt.Println(http.StatusServiceUnavailable)
		return
	}
}
