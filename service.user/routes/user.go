package routes

import (
	"net/http"

	. "github.com/agencyrevolution/go-microservices-example/models"

	"github.com/agencyrevolution/go-microservices-example/service.user/store"
	"github.com/gin-gonic/gin"
)

func AddRepoRoutes(s *store.Engine, r *gin.Engine) {
	r.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, s.State.Users)
	})

	r.GET("/users/:username", func(c *gin.Context) {
		var user *User

		for _, u := range s.State.Users {
			if u.Name == c.Param("username") {
				user = u
				break
			}
		}

		c.JSON(http.StatusOK, user)
	})

}
