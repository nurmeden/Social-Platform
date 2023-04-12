package controllers

import (
	"fmt"
	"net/http"
	"social-forum/internal/entity"
	"text/template"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case http.MethodGet:
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

	if password == Repeat_password {
		user := entity.User{
			Username:     username,
			Email:        email,
			PasswordHash: password,
		}
		err := h.services.User.CreateUser(user)
		if err != nil {
			fmt.Println(err)
		}
	}
	// case http.MethodPost:
	// 	http.Redirect(w, r, "/", http.StatusFound)
	// }
}

// func (h *Handler) UserByEmail(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/userbyemail" {
// 		fmt.Println("error from url", http.StatusNotFound)
// 		return
// 	}
// 	user, err := h.services.User.FindUserByEmail("nurmeden.02@gmail.com")
// 	if err != nil {
// 		fmt.Println("err in UserByEmail:", err)
// 		return
// 	}
// 	fmt.Println(user)
// }
