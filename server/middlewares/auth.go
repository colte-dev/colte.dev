package middlewares

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	accessToken := session.Get("access_token")

	if accessToken == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Set("access_token", accessToken)
}
