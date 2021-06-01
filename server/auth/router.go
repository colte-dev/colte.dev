package auth

import (
	"net/http"

	m "colte.dev/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(api *gin.RouterGroup) {
	router := api.Group("/auth")

	router.GET("/", m.AuthRequired, checkAuth)
	router.POST("/", authenticate)
	router.DELETE("/", logout)
}

func checkAuth(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func authenticate(c *gin.Context) {
	code, exists := c.GetQuery("code")

	if !exists || code == "" {
		c.Status(http.StatusBadRequest)
		return
	}
	token, err := GetAccessToken(code)
	errorMessage := err.Error()

	session := sessions.Default(c)

	if errorMessage != "" {
		c.Status(http.StatusUnauthorized)
		session.Delete("access_token")
	} else {
		session.Set("access_token", token)
		session.Options((sessions.Options{MaxAge: 288000, Path: "/"}))
		session.Save()
		c.Status(http.StatusNoContent)
	}
}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Status(http.StatusNoContent)
}
