package routes

import (
	"net/http"
	"social-forum/controllers"
)

func InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	// mux.HandleFunc("/refresh", Refresh)
	// mux.HandleFunc("/logout", Logout)
	// mux.HandleFunc("/login", Login)
	// mux.HandleFunc("/signUp", SignUp)
	// mux.HandleFunc("/Addpost", AddPost)
	// mux.HandleFunc("/posts", Posts)
	// mux.HandleFunc("/like", Like)
	fileServer := http.FileServer(http.Dir("./utils/"))
	mux.Handle("/utils/", http.StripPrefix("/utils/", fileServer))
	return mux
}
