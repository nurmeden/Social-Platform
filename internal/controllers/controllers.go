package controllers

import (
	"net/http"
	"social-forum/internal/usecase"
)

type Handler struct {
	services *usecase.Usecase
}

func NewHandler(services *usecase.Usecase) *Handler { // конструктор
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	mux.HandleFunc("/sign-up", h.SignUp)
	mux.HandleFunc("/userbyemail", h.UserByEmail)
	// mux.HandleFunc("/refresh", Refresh)
	// mux.HandleFunc("/logout", Logout)
	// mux.HandleFunc("/login", Login)
	// mux.HandleFunc("/Addpost", AddPost)
	// mux.HandleFunc("/posts", Posts)
	// mux.HandleFunc("/like", Like)

	fileServer := http.FileServer(http.Dir("./utils/"))

	mux.Handle("/utils/", http.StripPrefix("/utils/", fileServer))
	return mux
}
