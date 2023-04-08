package main

import (
	"fmt"
	"log"
	"net/http"
	"social-forum/routes"
)

func main() {
	log.Println("Запуск веб-сервера на http://localhost:8080/ ")
	err := http.ListenAndServe(":8080", routes.InitRoutes())
	if err != nil {
		fmt.Println(http.StatusServiceUnavailable)
		return
	}
}
