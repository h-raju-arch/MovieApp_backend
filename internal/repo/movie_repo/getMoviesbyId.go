package movierepo

import (
	"context"
	"fmt"

	"github.com/h-raju-arch/MoiveApp_backend/internal/model"
)

func (r MovieRepo) GetMovieBasebyId(ctx context.Context, id, lang string) (model.MovieResponse, error) {
	var res model.MovieResponse

	query := `SELECT 
	           m.id, COALESCE(mt.title,m.title) AS title, COALESCE(mt.overview,m.overview) AS overview,
			   to_char(m.release_date, 'YYYY-MM-DD') AS release_data,
			   ms.vote_average, ms.vote_count,
			   m.poster_path, m.backdrop_path,m.budget,m.revenue,m.homepage
			   FROM movies m
			   LEFT JOIN movie_translations mt ON mt.movie_id = m.id AND mt.language = $2
			   LEFT JOIN movie_stats ms ON ms.movie_id = m.id
			   WHERE m.id = $1;`

	err := r.db.QueryRowContext(ctx, query, id, lang).Scan(&res.ID,
		&res.Title,
		&res.Overview,
		&res.ReleaseDate,
		&res.VoteAverage,
		&res.VoteCount,
		&res.PosterPath,
		&res.BackdropPath,
		&res.Budget,
		&res.Revenue,
		&res.Homepage)
	if err != nil {
		return model.MovieResponse{}, fmt.Errorf("Query movie base: %w", err)
	}
	return res, nil
}

func (r MovieRepo) FetchGenres(ctx context.Context, id string) ([]string, error) {

	var res []string
	query := `SELECT g.name FROM genres g JOIN movie_genres mg ON g.id = mg.genre_id  WHERE mg.movie_id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)

	if err != nil {
		return nil, fmt.Errorf("Error Query FetchGenres: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, fmt.Errorf("Erro Fetch Genres row scan: %w", err)
		}
		res = append(res, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error Fetch Gneres row: %w", err)
	}
	return res, nil
}

func (r MovieRepo) FetchCredits(ctx context.Context, id string) ([]model.Credits_Response, error) {
	query := `SELECT p.name,p.known_for,c.credit_type
	          FROM people p JOIN credits c on p.id = c.person_id
			  WHERE c.movie_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return []model.Credits_Response{}, fmt.Errorf("Query FetchCredits : %w", err)
	}
	defer rows.Close()

	var resp []model.Credits_Response
	for rows.Next() {
		var temp model.Credits_Response
		err := rows.Scan(&temp.Name, &temp.Known_for, &temp.Credit_type)
		if err != nil {
			return []model.Credits_Response{}, fmt.Errorf("Error Credit rows scan: %w", err)
		}
		resp = append(resp, temp)
	}

	if err := rows.Err(); err != nil {
		return []model.Credits_Response{}, fmt.Errorf("Error Fetch Credits row: %w", err)
	}
	return resp, nil
}

func (r MovieRepo) FetchCompanies(ctx context.Context, id string) ([]string, error) {

	query := `SELECT c.name from companies c join movie_companies mc on
	         c.id = mc.company_id WHERE mc.movie_id = $1`

	rows, err := r.db.QueryContext(ctx, query, id)

	if err != nil {
		return nil, fmt.Errorf("Error Query Fetch Companies: %w", err)
	}

	defer rows.Close()
	var resp []string

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, fmt.Errorf("Error Fetch Comapanies scan: %w", err)
		}
		resp = append(resp, name)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error fetch companies rows: %w", err)
	}
	return resp, nil
}
