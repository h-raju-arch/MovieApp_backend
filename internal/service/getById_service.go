package service

import (
	"context"
	"fmt"

	"github.com/h-raju-arch/MoiveApp_backend/internal/model"
)

//getMoviebyId

func (r Movie_services) GetMovieById(ctx context.Context, id, lang string, appendtoresponse []string) (model.MovieResponse, error) {
	movie, err := r.repo.GetMovieBasebyId(ctx, id, lang)

	if err != nil {
		return model.MovieResponse{}, fmt.Errorf("service: Get base movie: %w", err)
	}

	var genres []string
	var companies []string
	var credits []model.Credits_Response

	for _, append := range appendtoresponse {
		switch append {
		case "genre":
			genres, err = r.repo.FetchGenres(ctx, id)
			if err != nil {
				return model.MovieResponse{}, fmt.Errorf("fetch Genres: %w", err)
			}
		case "companies":
			companies, err = r.repo.FetchCompanies(ctx, id)
			if err != nil {
				return model.MovieResponse{}, fmt.Errorf("companies Genres: %w", err)
			}

		case "credits":
			credits, err = r.repo.FetchCredits(ctx, id)
			if err != nil {
				return model.MovieResponse{}, fmt.Errorf("fetch credits: %w", err)
			}
		}
	}

	res := model.MovieResponse{
		ID:           movie.ID,
		Title:        movie.Title,
		Overview:     movie.Overview,
		ReleaseDate:  movie.ReleaseDate,
		VoteAverage:  movie.VoteAverage,
		VoteCount:    movie.VoteCount,
		PosterPath:   movie.PosterPath,
		BackdropPath: movie.BackdropPath,
		Budget:       movie.Budget,
		Revenue:      movie.Revenue,
		Homepage:     movie.Homepage,
	}

	if contains(appendtoresponse, "genre") {
		res.Genres = genres
	}
	if contains(appendtoresponse, "companies") {
		res.ProductionCompanies = companies
	}
	if contains(appendtoresponse, "credits") {
		res.Credits = credits
	}
	return res, nil
}

func contains(list []string, s string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}
