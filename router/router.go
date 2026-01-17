package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type routerStruct struct {
	engine *gin.Engine
}

func (r *routerStruct) EcommerceHealthCheck(rg *gin.RouterGroup) {
	// Grouping routes under /ecommerce
	orderRouteGrouping := rg.Group("/ecommerce")
	orderRouteGrouping.Use(CORSMiddleware())

	// healthCheckRoutes is expected to be defined in another file within package router
	for _, route := range healthCheckRoutes {
		switch route.Method {
		case http.MethodGet:
			orderRouteGrouping.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			orderRouteGrouping.POST(route.Pattern, route.HandlerFunc)
		case http.MethodOptions:
			orderRouteGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			orderRouteGrouping.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			orderRouteGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			orderRouteGrouping.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"result": "Specify a valid HTTP method with this router",
				})
			})
		}
	}
}

func ClientRoutes() {
	r := gin.Default()

	// Initialize the routerStruct
	routerObj := &routerStruct{
		engine: r,
	}

	// Define API version group
	v1 := r.Group("/api/v1")

	// Call the method to register routes
	routerObj.EcommerceHealthCheck(v1)

	// Fetch port from env or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

// CORSMiddleware handles Cross-Origin Resource Sharing
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}