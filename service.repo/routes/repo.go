package routes

import (
	"net/http"

	. "github.com/agencyrevolution/go-microservices-example/models"

	"github.com/agencyrevolution/go-microservices-example/service.repo/store"
	"github.com/gin-gonic/gin"
)

func AddRepoRoutes(s *store.Engine, r *gin.Engine) {
	r.GET("/users/:username/repos", func(c *gin.Context) {
		var repos []*Repo

		for _, repo := range s.State.Repos {
			if repo.RepoOwner.Username == c.Param("username") {
				repos = append(repos, repo)
			}
		}

		c.JSON(http.StatusOK, repos)
	})

	r.GET("/users/:username/repos/:reponame", func(c *gin.Context) {
		var repo *Repo

		for _, rp := range s.State.Repos {
			if rp.RepoOwner.Username == c.Param("username") &&
				rp.Name == c.Param("reponame") {
				repo = rp
				break
			}
		}

		c.JSON(http.StatusOK, repo)
	})
}
