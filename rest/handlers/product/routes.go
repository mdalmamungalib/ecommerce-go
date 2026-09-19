package product

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manger) {

	mux.Handle(
		"GET /products",
		middleware.Cors(
			middleware.Preflight(
				middleware.Logger(
					http.HandlerFunc(h.GetProducts),
				),
			),
		),
	)

	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProduct),
		),
	)

	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
			h.middlewares.AuthenticateJWT,
		),
	)

	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.middlewares.AuthenticateJWT,
		),
	)
}
