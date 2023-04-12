package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/sign-in" {
			fmt.Println("error from url", http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles("./templates/sign-in.html")
		if err != nil {
			fmt.Println(http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			fmt.Println(http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		username := r.FormValue("uname")
		password := r.FormValue("psw")

		user, err := h.services.User.FindUserByUsername(username, password)
		if err != nil {
			log.Printf("Username doesn't exist")
			fmt.Println(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "sessions",
			Value: user.Username,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
