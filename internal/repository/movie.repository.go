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
			m.id, 
			m.title, 
			m.poster_url, 
			m.release_date, 
			STRING_AGG(DISTINCT g.name, ', ') AS genres
		FROM movies m 
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id 
		LEFT JOIN genres g ON mg.genre_id = g.id 
		WHERE m.release_date > CURRENT_DATE 
		GROUP BY m.id, m.title, m.poster_url, m.release_date 
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
			&movie.Id,
			&movie.Title,
			&movie.PosterUrl,
			&movie.ReleaseDate,
			&movie.GenresName,
		)
		if err != nil {
			log.Println("Scan error:", err.Error())
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
