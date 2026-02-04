package router

import (
	"net/http"
	"threadStocks/core"
)

/**
* Applique les Middleware de haut en bas
* ```go
* mux.Handle("/api/users", chain(http.HandlerFunc(usersHandler), authMiddleware, loggingMiddleware))
* ```
 */
func chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

func Router(s *http.ServeMux, a *core.App) {
	s.HandleFunc("/login", a.Controller.Auth.HandleLogin)
	s.HandleFunc("/register", a.Controller.Auth.HandleRegister)
	s.Handle("/users/me", chain(
		http.HandlerFunc(a.Controller.User.Me),
		core.AuthMiddleware))
	s.Handle("/threads", chain(
		http.HandlerFunc(a.Controller.Thread.GetAllThreadByUser),
		core.AuthMiddleware))
	s.Handle("/threads/{id}", chain(
		http.HandlerFunc(a.Controller.Thread.GetThread),
		core.AuthMiddleware))
	s.Handle("/threads/create", chain(
		http.HandlerFunc(a.Controller.Thread.CreateThread),
		core.AuthMiddleware))
	s.Handle("/threads/update/{id}", chain(
		http.HandlerFunc(a.Controller.Thread.UpdateThread),
		core.AuthMiddleware))
	s.Handle("/threads/delete/{id}", chain(
		http.HandlerFunc(a.Controller.Thread.DeleteThread),
		core.AuthMiddleware))
	s.Handle("/threads/delete", chain(
		http.HandlerFunc(a.Controller.Thread.DeleteMultipleThread),
		core.AuthMiddleware))
	s.Handle("/threads/update", chain(
		http.HandlerFunc(a.Controller.Thread.UpdateMultipleThread),
		core.AuthMiddleware))
}