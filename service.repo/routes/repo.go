package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/quangbuule/go-microservices-example/service.repo/store"
)

func AddRepoRoutes(s *store.Engine, r *gin.Engine) {
	r.GET("/users/:username/repos", func(c *gin.Context) {
		var repos []*store.Repo

		for _, repo := range s.State.Repos {
			if repo.RepoOwner.Username == c.Param("username") {
				repos = append(repos, repo)
			}
		}

		c.JSON(http.StatusOK, repos)
	})

	r.GET("/users/:username/repos/:reponame", func(c *gin.Context) {
		var repo *store.Repo

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