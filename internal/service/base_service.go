package service

import (
	"context"

	"github.com/h-raju-arch/MoiveApp_backend/internal/model"
	movierepo "github.com/h-raju-arch/MoiveApp_backend/internal/repo/movie_repo"
)

type MovieService interface {
	GetMovieById(ctx context.Context, id, lang string, appendtoresponse []string) (model.MovieResponse, error)
}

type Movieservices struct {
	repo movierepo.MovieRepo
}

func New_Movie_Service(r movierepo.MovieRepo) *Movieservices {
	return &Movieservices{repo: r}
}
