package backend

import (
	"net/http"

	"github.com/Dreck2003/api/backend/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateRoutes() *chi.Mux {
	rootRouter := chi.NewRouter()
	rootRouter.Use(middleware.Logger)
	rootRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	routes.CreateAllRoutes(rootRouter)

	return rootRouter
}
