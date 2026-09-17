package rest

// import (
// 	middleware "ecommerce/rest/middlewares"
// 	"net/http"
// )

// func (h *Handler) initRoutes(mux *http.ServeMux, manager *middleware.Manger) {

// 	mux.Handle(
// 		"GET /products",
// 		middleware.Cors(
// 			middleware.Preflight(
// 				middleware.Logger(
// 					http.HandlerFunc(handlers.GetProducts),
// 				),
// 			),
// 		),
// 	)

// 	mux.Handle(
// 		"POST /products",
// 		manager.With(
// 			http.HandlerFunc(handlers.CreateProduct),
// 			middleware.AuthenticateJWT,
// 		),
// 	)

// 	mux.Handle(
// 		"GET /products/{id}",
// 		manager.With(
// 			http.HandlerFunc(handlers.GetProduct),
// 		),
// 	)

// 	mux.Handle(
// 		"PUT /products/{id}",
// 		manager.With(
// 			http.HandlerFunc(handlers.UpdateProduct),
// 			middleware.AuthenticateJWT,
// 		),
// 	)

// 	mux.Handle(
// 		"DELETE /products/{id}",
// 		manager.With(
// 			http.HandlerFunc(handlers.DeleteProduct),
// 			middleware.AuthenticateJWT,
// 		),
// 	)

// 	mux.Handle(
// 		"POST /users",
// 		manager.With(
// 			http.HandlerFunc(handlers.CreateUser),
// 		),
// 	)

// 	mux.Handle(
// 		"POST /login",
// 		manager.With(
// 			http.HandlerFunc(handlers.Login),
// 		),
// 	)
// }
