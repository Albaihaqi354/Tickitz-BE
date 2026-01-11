package repository

import (
	"context"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminRepository struct {
	db *pgxpool.Pool
}

func NewAdminRepository(db *pgxpool.Pool) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

func (a AdminRepository) GetAllMovieAdmin(ctx context.Context) ([]model.MovieDetail, error) {
	sqlStr := `
		SELECT 
			m.id,
			m.title,
			COALESCE(m.synopsis, '') AS synopsis,
			COALESCE(m.duration, 0) AS duration,
			m.release_date,
			COALESCE(d.name, '') AS director,
			COALESCE(STRING_AGG(DISTINCT a.name, ', '), '') AS "cast",
			COALESCE(m.poster_url, '') AS poster_url,
			COALESCE(m.backdrop_url, '') AS backdrop_url,
			COALESCE(m.popularity_score, 0) AS popularity_score,
			m.created_at,
			m.updated_at,
			COALESCE(STRING_AGG(DISTINCT g.name, ', '), '') AS genre_name,
			COUNT(DISTINCT s.id) AS schedule_count
		FROM movies m
		LEFT JOIN directors d ON m.director_id = d.id
		LEFT JOIN movie_casts mc ON m.id = mc.movie_id
		LEFT JOIN actors a ON mc.actor_id = a.id
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
		LEFT JOIN schedules s ON s.movie_id = m.id
		GROUP BY m.id, d.name
		ORDER BY m.created_at DESC;`

	rows, err := a.db.Query(ctx, sqlStr)
	if err != nil {
		log.Println("Query Error:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var movies []model.MovieDetail
	for rows.Next() {
		var movie model.MovieDetail
		err := rows.Scan(
			&movie.Id,
			&movie.Title,
			&movie.Synopsis,
			&movie.Duration,
			&movie.ReleaseDate,
			&movie.Director,
			&movie.Cast,
			&movie.PosterUrl,
			&movie.BackdropUrl,
			&movie.PopularityScore,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&movie.GenresName,
			&movie.ScheduleCount,
		)
		if err != nil {
			log.Println("Scan Error:", err.Error())
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (a AdminRepository) DeleteMovieAdmin(ctx context.Context, movieId int) error {
	sqlStr := `
		DELETE FROM movies
		WHERE id = $1;
	`

	_, err := a.db.Exec(ctx, sqlStr, movieId)
	if err != nil {
		log.Println("Query Error:", err.Error())
		return err
	}

	return nil
}
