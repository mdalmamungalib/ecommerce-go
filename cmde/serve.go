package cmde

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve() {
	cnf := config.GetConfig()

	
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

	fmt.Println("Server running on :", cnf.HttpPort)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
