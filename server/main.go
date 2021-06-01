package main

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"colte.dev/auth"
	c "colte.dev/common"
	"colte.dev/discussions"
	"colte.dev/env"
	"colte.dev/middlewares"
)

func main() {
	env.Init()

	c.InitRedis()

	// create gin instance
	app := gin.Default()

	// session middleware
	store := cookie.NewStore([]byte(env.COOKIE_SECRET))
	app.Use(sessions.Sessions("colte-session", store))

	// cors middleware
	app.Use(middlewares.CORSMiddleware())

	// register routes
	api := app.Group("/api")
	discussions.RegisterRoute(api)
	auth.RegisterRoute(api)

	// health check endpoint
	api.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// run app
	app.Run(":" + os.Getenv("PORT"))
}
