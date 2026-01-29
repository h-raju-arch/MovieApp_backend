package service

import (
	"context"

	"github.com/h-raju-arch/MoiveApp_backend/internal/model"
	movierepo "github.com/h-raju-arch/MoiveApp_backend/internal/repo/movie_repo"
)

type Movie_Service interface {
	GetMovieById(ctx context.Context, id, lang string, appendtoresponse []string) (model.MovieResponse, error)
}

type Movie_services struct {
	repo movierepo.Movie_repo
}

func New_Movie_Service(r movierepo.Movie_repo) *Movie_services {
	return &Movie_services{repo: r}
}
