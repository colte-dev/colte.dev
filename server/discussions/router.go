package discussions

import (
	"net/http"
	"strconv"

	m "colte.dev/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(api *gin.RouterGroup) {
	router := api.Group("/discussions")
	router.Use(m.AuthRequired)

	router.GET("/", ControllerGetDiscussions)
	router.GET("/:number", ControllerGetByNumber)
}

func ControllerGetByNumber(c *gin.Context) {
	number, err := strconv.Atoi(c.Params.ByName("number"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	accessToken, _ := c.Get("access_token")
	discussion, err := GetDiscussionByNumber(number, accessToken.(string))

	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, discussion)
	}
}

func ControllerGetDiscussions(c *gin.Context) {
	username, exists := c.GetQuery("username")
	accessToken, _ := c.Get("access_token")

	var discussions []Discussion
	var err error

	if exists && username != "" {
		discussions, err = GetUserDiscussions(username, accessToken.(string))
	} else {
		discussions, err = GetLastDiscussions(accessToken.(string))
	}

	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, discussions)
	}
}
