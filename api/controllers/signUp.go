package controllers

import (
	"fmt"
	"net/http"
	"social-forum/models"
	"text/template"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/sign-up" {
			fmt.Println("error from url", http.StatusNotFound)
			return
		}
		tmpl, err := template.ParseFiles("./templates/signUp.html")
		if err != nil {
			fmt.Println(http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(http.StatusInternalServerError)
			return
		}
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("psw")
		Repeat_password := r.FormValue("psw-repeat")
		fmt.Printf("username: %v\n", username)
		if password == Repeat_password {
			var user models.User
			user = models.User{
				Email:    email,
				Name:     username,
				Password: password,
			}
			fmt.Printf("user: %v\n", user)
			err = h.services.CreateUser(user)
			if err != nil {
				fmt.Println(err)
			}
		}
	case http.MethodPost:
		http.Redirect(w, r, "/", 302)
	}
}
