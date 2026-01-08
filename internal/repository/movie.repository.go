package repository

import (
	"context"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieRepository struct {
	db *pgxpool.Pool
}

func NewMoviesRepository(db *pgxpool.Pool) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (m MovieRepository) GetUpcomingMovie(ctx context.Context) ([]model.MovieDetail, error) {
	sqlStr := `
		SELECT 
			m.id, m.title, m.synopsis, m.duration, m.release_date, 
			m.director_id, d.name AS director_name, 
			STRING_AGG(DISTINCT a.name, ', ') AS "cast", 
			m.poster_url, m.backdrop_url, m.popularity_score, 
			STRING_AGG(DISTINCT g.name, ', ') AS genre_name,
			m.created_at, m.updated_at
		FROM movies m 
		LEFT JOIN directors d ON m.director_id = d.id 
		LEFT JOIN movie_casts mc ON m.id = mc.movie_id 
		LEFT JOIN actors a ON mc.actor_id = a.id 
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id 
		LEFT JOIN genres g ON mg.genre_id = g.id 
		WHERE m.release_date > CURRENT_DATE 
		GROUP BY m.id, d.name 
		ORDER BY m.release_date ASC;`
	rows, err := m.db.Query(ctx, sqlStr)
	if err != nil {
		log.Println("Query error:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var movies []model.MovieDetail
	for rows.Next() {
		var movie model.MovieDetail
		err := rows.Scan(
			&movie.Id, &movie.Title, &movie.Synopsis, &movie.Duration, &movie.ReleaseDate,
			&movie.Director.Id, &movie.Director.Name, &movie.Cast, &movie.PosterUrl, &movie.BackDropUrl,
			&movie.PopularityScore, &movie.GenresName, &movie.CreatedAt, &movie.UpdatedAt,
		)
		if err != nil {
			log.Println("Scan error:", err.Error())
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
