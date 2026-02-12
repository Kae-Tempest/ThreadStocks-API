package router

import (
	"net/http"
	"threadStocks/core"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
	s.Handle("/login", otelhttp.NewHandler(http.HandlerFunc(a.Controller.Auth.HandleLogin), "HandleLogin"))
	s.Handle("/register", otelhttp.NewHandler(http.HandlerFunc(a.Controller.Auth.HandleRegister), "HandleRegister"))
	s.Handle("/users/me", otelhttp.NewHandler(chain(
		http.HandlerFunc(a.Controller.User.Me),
		core.AuthMiddleware), "Me"))
	s.Handle("/threads", otelhttp.NewHandler(chain(
		http.HandlerFunc(a.Controller.Thread.GetAllThreadByUser),
		core.AuthMiddleware), "GetAllThreadByUser"))
	s.Handle("/threads/{id}", otelhttp.NewHandler(chain(
		http.HandlerFunc(a.Controller.Thread.GetThread),
		core.AuthMiddleware), "GetThread"))
	s.Handle("/threads/create", otelhttp.NewHandler(chain(
		http.HandlerFunc(a.Controller.Thread.CreateThread),
		core.AuthMiddleware), "CreateThread"))
	s.Handle("/threads/update/{id}", otelhttp.NewHandler(chain(
		http.HandlerFunc(a.Controller.Thread.UpdateThread),
		core.AuthMiddleware), "UpdateThread"))
	s.Handle("/threads/delete/{id}", otelhttp.NewHandler(chain(
		http.HandlerFunc(a.Controller.Thread.DeleteThread),
		core.AuthMiddleware), "DeleteThread"))
	s.Handle("/threads/delete", otelhttp.NewHandler(chain(
		http.HandlerFunc(a.Controller.Thread.DeleteMultipleThread),
		core.AuthMiddleware), "DeleteMultipleThread"))
	s.Handle("/threads/update", otelhttp.NewHandler(chain(
		http.HandlerFunc(a.Controller.Thread.UpdateMultipleThread),
		core.AuthMiddleware), "UpdateMultipleThread"))
}