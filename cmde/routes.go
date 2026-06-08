package cmde

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manger) {
	mux.Handle(
		"GET /rahim",
		manager.With(
			http.HandlerFunc(handlers.Test),
			middleware.Arekta,
		),
	)

	mux.Handle(
		"GET /mamun",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)

	mux.Handle(
		"GET /products",
		middleware.Cors(
			middleware.Preflight(
				middleware.Logger(
					http.HandlerFunc(handlers.GetProducts),
				),
			),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductByID),
		),
	)
}
