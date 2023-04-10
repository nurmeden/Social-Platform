package controllers

import (
	"fmt"
	"net/http"
	"social-forum/internal/entity"
	"text/template"
)

var users entity.User

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.URL.Path)
		if r.URL.Path != "/" {
			fmt.Println("error from url", http.StatusNotFound)
			return
		}
	}
	tmpl, err := template.ParseFiles("./templates/home.html")
	if err != nil {
		fmt.Println("error parsing file ", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, users)
	if err != nil {
		fmt.Println("error from executing: ", http.StatusInternalServerError)
		return
	}
	fmt.Println(r.Form)
}
