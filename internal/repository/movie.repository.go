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

func (m MovieRepository) GetPopularMovie(ctx context.Context) ([]model.MovieDetail, error) {
	sqlStr := `
		SELECT 
			m.id,
			m.title,
			m.poster_url,
			m.popularity_score,
			STRING_AGG(DISTINCT g.name, ', ') AS genre_name
		FROM movies m
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
		GROUP BY m.id, m.title, m.poster_url, m.popularity_score
		ORDER BY m.popularity_score DESC;`
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
			&movie.PopularityScore,
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

func (m MovieRepository) GetMovieWithFilter(ctx context.Context, search *string, genreId *int) ([]model.MovieDetail, error) {
	sqlStr := `
		SELECT 
			m.id,
			m.title,
			m.poster_url,
			STRING_AGG(DISTINCT g.name, ', ') AS genre_name
		FROM movies m
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
		WHERE 
			($1::TEXT IS NULL OR m.title ILIKE '%' || $1 || '%') AND
			($2::INTEGER IS NULL OR m.id IN (
				SELECT movie_id FROM movie_genres WHERE genre_id = $2
			))
		GROUP BY m.id, m.title, m.poster_url, m.release_date
		ORDER BY m.release_date DESC;`

	rows, err := m.db.Query(ctx, sqlStr, search, genreId)
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

func (m MovieRepository) GetMovieDetail(ctx context.Context, idDetail *int) ([]model.MovieDetail, error) {
	sqlStr := `
		SELECT
			m.id,
			m.title,
			m.synopsis,
			m.duration,
			m.release_date,
			d.name AS director,
			STRING_AGG(DISTINCT a.name, ', ') AS "cast",
			m.poster_url,
			m.backdrop_url,
			m.popularity_score,
			STRING_AGG(DISTINCT g.name, ', ') AS genre_name
		FROM movies m
		LEFT JOIN directors d ON m.director_id = d.id
		LEFT JOIN movie_casts mc ON m.id = mc.movie_id
		LEFT JOIN actors a ON mc.actor_id = a.id
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
		WHERE m.id = $1
		GROUP BY m.id, d.name;`

	rows, err := m.db.Query(ctx, sqlStr, idDetail)
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
			&movie.Synopsis,
			&movie.Duration,
			&movie.ReleaseDate,
			&movie.Director,
			&movie.Cast,
			&movie.PosterUrl,
			&movie.BackDropUrl,
			&movie.PopularityScore,
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
