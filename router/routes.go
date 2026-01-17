import (
    "net/http"
    "ecommerce/constant"
    "ecommerce/controller"
)

var healthCheckRouter = Router{
    {
        Name:        "Health check",
        Method:      http.MethodGet,
        Pattern:     constant.HealthCheckRoute,
        HandlerFunc: controller.HealthCheck,
    },
}