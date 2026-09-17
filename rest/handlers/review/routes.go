package review

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manger) {
	mux.Handle(
		"GET /reviews",
		middleware.Cors(
			middleware.Preflight(
				middleware.Logger(
					http.HandlerFunc(h.GetReviews),
				),
			),
		),
	)
}
