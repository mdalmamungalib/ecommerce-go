package cmde

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	
	manager := middleware.NewManger()

	// globalRouter := middleware.CorsWithPreflight(mux)
	// wrappedMux := manager.WrapMux(mux,
	// 	middleware.Logger,
	// 	middleware.Hudai,
	// 	middleware.CorsWithPreflight,
	// )

	manager.Use(
		middleware.Preflight,
		middleware.Cors, 
		middleware.Logger,
	)
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
