package repository

import (
	"context"
	"log"

	"github.com/Albaihaqi354/Tickitz-BE/internal/dto"
	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) GetProfile(ctx context.Context, userId int) (model.User, error) {
	sqlStr := `
		SELECT 
			id, 
			email, 
			COALESCE(first_name, ''), 
			COALESCE(last_name, ''), 
			COALESCE(phone_number, ''), 
			COALESCE(profile_image, ''), 
			loyalty_points, 
			role,
			created_at
		FROM users
		WHERE id = $1;`

	row := u.db.QueryRow(ctx, sqlStr, userId)

	var user model.User
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.PhoneNumber,
		&user.ProfileImage,
		&user.LoyaltyPoints,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		log.Println("Scan error:", err.Error())
		return model.User{}, err
	}

	return user, nil
}

func (u UserRepository) GetHistory(ctx context.Context, userId int) ([]model.GetHistory, error) {
	sqlStr := `
		SELECT 
			o.id AS order_id,
			o.booking_code,
			o.total_price,
			o.payment_status,
			o.created_at AS order_date,
			m.id AS movie_id,
			m.title AS movie_title,
			m.poster_url AS movie_poster,
			c.name AS cinema_name,
			c.logo_url AS cinema_logo,
			s.show_date,
			s.show_time,
			COUNT(od.id) AS ticket_count
		FROM orders o
		INNER JOIN schedules s ON o.schedule_id = s.id
		INNER JOIN movies m ON s.movie_id = m.id
		INNER JOIN cinemas c ON s.cinema_id = c.id
		LEFT JOIN order_details od ON od.order_id = o.id
		WHERE o.user_id = $1
		GROUP BY o.id, m.id, c.id, s.id
		ORDER BY o.created_at DESC;`

	rows, err := u.db.Query(ctx, sqlStr, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var histories []model.GetHistory
	for rows.Next() {
		var h model.GetHistory
		err := rows.Scan(
			&h.Id,
			&h.BookingCode,
			&h.TotalPrice,
			&h.PaymentStatus,
			&h.CreatedAt,
			&h.MovieId,
			&h.Title,
			&h.PosterUrl,
			&h.CinemaName,
			&h.CinemaLogo,
			&h.ShowDate,
			&h.ShowTime,
			&h.TicketCount,
		)
		if err != nil {
			log.Println("Scan error:", err.Error())
			return nil, err
		}
		histories = append(histories, h)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return histories, nil
}

func (u UserRepository) GetPasswordById(ctx context.Context, userId int) (string, error) {
	sqlStr := "SELECT password FROM users WHERE id = $1"
	var password string
	err := u.db.QueryRow(ctx, sqlStr, userId).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}

func (u UserRepository) UpdatePassword(ctx context.Context, userId int, hashedPassword string) error {
	sqlStr := "UPDATE users SET password = $1, updated_at = now() WHERE id = $2"
	_, err := u.db.Exec(ctx, sqlStr, hashedPassword, userId)
	return err
}

func (u UserRepository) UpdateProfile(ctx context.Context, userId int, req dto.UpdateProfileRequest) (model.User, error) {
	sqlStr := `
		UPDATE users
		SET
			first_name = COALESCE($1, first_name),
			last_name = COALESCE($2, last_name),
			phone_number = COALESCE($3, phone_number),
			profile_image = COALESCE($4, profile_image),
			updated_at = NOW()
		WHERE id = $5
		RETURNING id, COALESCE(first_name, ''), COALESCE(last_name, ''), COALESCE(phone_number, ''), COALESCE(profile_image, '');
	`

	var m model.User
	err := u.db.QueryRow(ctx, sqlStr,
		req.FirstName,
		req.LastName,
		req.PhoneNumber,
		req.ProfileImage,
		userId,
	).Scan(
		&m.Id,
		&m.FirstName,
		&m.LastName,
		&m.PhoneNumber,
		&m.ProfileImage,
	)

	if err != nil {
		log.Println("Update Error", err.Error())
		return model.User{}, err
	}
	return m, nil
}
