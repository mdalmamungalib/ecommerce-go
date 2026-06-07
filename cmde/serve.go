package cmde

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManger()

	manager.Use(middleware.Logger, middleware.Hudai, )

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	fmt.Println("Server running on :8080")

	globalRouter := middleware.CorsWithPreflight(mux)

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
