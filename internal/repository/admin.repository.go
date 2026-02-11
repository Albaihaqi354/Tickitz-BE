package repository

import (
	"context"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
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

func (a AdminRepository) UpdateMovieAdmin(ctx context.Context, id int, req dto.UpdateMovieRequest) (model.Movie, error) {
	sqlStr := `
		UPDATE movies
		SET 
			title = COALESCE($1, title),
			synopsis = COALESCE($2, synopsis),
			duration = COALESCE($3, duration),
			release_date = COALESCE($4, release_date),
			director_id = COALESCE($5, director_id),
			poster_url = COALESCE($6, poster_url),
			backdrop_url = COALESCE($7, backdrop_url),
			popularity_score = COALESCE($8, popularity_score),
			updated_at = NOW()
		WHERE id = $9
		RETURNING id, title, synopsis, duration, release_date, director_id, poster_url, backdrop_url, popularity_score;`

	var m model.Movie
	err := a.db.QueryRow(ctx, sqlStr,
		req.Title,
		req.Synopsis,
		req.Duration,
		req.ReleaseDate,
		req.DirectorId,
		req.PosterUrl,
		req.BackdropUrl,
		req.PopularityScore,
		id,
	).Scan(
		&m.Id,
		&m.Title,
		&m.Synopsis,
		&m.Duration,
		&m.ReleaseDate,
		&m.DirectorId,
		&m.PosterUrl,
		&m.BackdropUrl,
		&m.PopularityScore,
	)

	if err != nil {
		log.Println("Update Error:", err.Error())
		return model.Movie{}, err
	}

	return m, nil
}

func (a AdminRepository) CreateMovieAdmin(ctx context.Context, req dto.CreateMovieRequest) (model.Movie, error) {
	tx, err := a.db.Begin(ctx)
	if err != nil {
		log.Println("Transaction Begin Error:", err.Error())
		return model.Movie{}, err
	}
	defer tx.Rollback(ctx)

	sqlStr := `
		INSERT INTO movies (
			title, 
			synopsis, 
			duration, 
			release_date, 
			director_id, 
			poster_url, 
			backdrop_url, 
			popularity_score, 
			created_at, 
			updated_at
		) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
		RETURNING id, title, synopsis, duration, release_date, director_id, poster_url, backdrop_url, popularity_score;`

	var m model.Movie
	err = tx.QueryRow(ctx, sqlStr,
		req.Title,
		req.Synopsis,
		req.Duration,
		req.ReleaseDate,
		req.DirectorId,
		req.PosterUrl,
		req.BackdropUrl,
		req.PopularityScore,
	).Scan(
		&m.Id,
		&m.Title,
		&m.Synopsis,
		&m.Duration,
		&m.ReleaseDate,
		&m.DirectorId,
		&m.PosterUrl,
		&m.BackdropUrl,
		&m.PopularityScore,
	)

	if err != nil {
		log.Println("Insert Movie Error:", err.Error())
		return model.Movie{}, err
	}

	if len(req.Genres) > 0 {
		for _, genreId := range req.Genres {
			genreSql := `INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)`
			_, err := tx.Exec(ctx, genreSql, m.Id, genreId)
			if err != nil {
				log.Println("Insert Genre Error:", err.Error())
				return model.Movie{}, err
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Println("Transaction Commit Error:", err.Error())
		return model.Movie{}, err
	}

	return m, nil
}
