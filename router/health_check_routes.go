package router

import (
	"ecommerce/controller"
	"net/http"
)

var healthCheckRoutes = Router{
	Route{
		Name:        "HealthCheck",
		Method:      http.MethodGet,
		Pattern:     "/health",
		HandlerFunc: controller.HealthCheck,
	},
}
