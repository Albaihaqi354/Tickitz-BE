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

func (u UserRepository) CreateNewUser(ctx context.Context, newUser dto.NewUser) (model.User, error) {
	sql := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id, email, password, role"
	values := []any{newUser.Email, newUser.Password}

	row := u.db.QueryRow(ctx, sql, values...)
	var user model.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepository) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	sql := "SELECT id, email, password, role FROM users WHERE email = $1"

	row := u.db.QueryRow(ctx, sql, email)
	var user model.User
	if err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (u UserRepository) GetProfile(ctx context.Context, userId int) (model.User, error) {
	sqlStr := `
		SELECT 
			id, 
			email, 
			first_name, 
			last_name, 
			phone_number, 
			profile_image, 
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
