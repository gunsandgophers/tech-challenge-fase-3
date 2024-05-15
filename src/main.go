package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"tech-challenge-fase-1/internal/adapter/driver/routers"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// config cors
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowHeaders = []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"}

	r.Use(cors.New(config))
	r = routers.RegisterRouters(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
