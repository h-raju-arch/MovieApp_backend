package httptransport

import (
	"github.com/gin-gonic/gin"
	"github.com/h-raju-arch/MoiveApp_backend/internal/service"
)

func NewRouter(movie_svc service.Movie_Service) *gin.Engine {
	router := gin.Default()
	h := New_Movie_Handler(movie_svc)

	api := router.Group("/api")
	{
		api.GET("/movie/", h.GetMovies)
	}
	return router
}
