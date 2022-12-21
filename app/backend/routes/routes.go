package routes

import (
	"github.com/Dreck2003/api/backend/routes/api"
	"github.com/Dreck2003/api/backend/types"
	"github.com/go-chi/chi/v5"
)

func CreateAllRoutes(rootRouter *chi.Mux) {

	rootRouter.Route("/emails", func(r chi.Router) {
		for _, s := range api.EmailRoutes {
			switch s.TypeMethod {
			case types.Get:
				r.Get(s.Pattern, s.Router)
			case types.Post:
				r.Post(s.Pattern, s.Router)
			case types.Delete:
				r.Post(s.Pattern, s.Router)
			case types.Put:
				r.Post(s.Pattern, s.Router)
			}
		}
	})
}
