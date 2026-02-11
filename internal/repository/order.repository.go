package repository

import (
	"context"
	"log"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type OrderRepo interface {
	GetSchedules(ctx context.Context, db DBTX, movieId int, showDate *string, city *string) ([]model.GetSchedules, error)
	InsertOrder(ctx context.Context, db DBTX, order model.Order) (int, string, time.Time, error)
	InsertOrderDetail(ctx context.Context, db DBTX, orderId int, seatId int) error
	GetSeatsByScheduleID(ctx context.Context, db DBTX, scheduleId int) ([]model.Seat, error)
	GetPriceFromSchedule(ctx context.Context, db DBTX, scheduleId int) (int, error)
	UpdatePaymentStatus(ctx context.Context, db DBTX, orderId int, status string) error
}

type OrderRepository struct{}

func NewOrdersRepository() *OrderRepository {
	return &OrderRepository{}
}

func (o OrderRepository) GetSchedules(ctx context.Context, db DBTX, movieId int, showDate *string, city *string) ([]model.GetSchedules, error) {
	sqlStr := `
		SELECT 
			s.id,
			s.show_date,
			s.show_time,
			s.price,
			c.id AS cinema_id,
			c.name AS cinema_name,
			c.logo_url AS cinema_logo,
			c.location AS cinema_location,
			ci.name AS cinema_city
		FROM schedules s
		INNER JOIN cinemas c ON s.cinema_id = c.id
		INNER JOIN cities ci ON c.city_id = ci.id
		WHERE s.movie_id = $1
			AND ($2::DATE IS NULL OR s.show_date = $2)
			AND ($3::VARCHAR IS NULL OR ci.name = $3)
		ORDER BY s.show_date, s.show_time;
	`

	rows, err := db.Query(ctx, sqlStr, movieId, showDate, city)
	if err != nil {
		log.Println("Query error:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var schedules []model.GetSchedules
	for rows.Next() {
		var s model.GetSchedules
		err := rows.Scan(
			&s.Id,
			&s.ShowDate,
			&s.ShowTime,
			&s.Price,
			&s.CinemaId,
			&s.CinemaName,
			&s.CinemaLogo,
			&s.CinemaLocation,
			&s.CinemaCity,
		)
		if err != nil {
			log.Println("Scan error:", err.Error())
			return nil, err
		}
		schedules = append(schedules, s)
	}
	return schedules, nil
}

func (o OrderRepository) InsertOrder(ctx context.Context, db DBTX, order model.Order) (int, string, time.Time, error) {
	sqlOrder := `
		INSERT INTO orders (user_id, schedule_id, total_price, payment_status, booking_code)
		VALUES ($1, $2, $3, $4, md5(random()::text))
		RETURNING id, booking_code, created_at`

	var id int
	var bookingCode string
	var createdAt time.Time

	err := db.QueryRow(ctx, sqlOrder, order.UserId, order.ScheduleId, order.TotalPrice, order.PaymentStatus).Scan(&id, &bookingCode, &createdAt)
	if err != nil {
		log.Println("InsertOrder Error:", err.Error())
		return 0, "", time.Time{}, err
	}

	return id, bookingCode, createdAt, nil
}

func (o OrderRepository) InsertOrderDetail(ctx context.Context, db DBTX, orderId int, seatId int) error {
	sqlOrderDetail := `INSERT INTO order_details (order_id, seat_id) VALUES ($1, $2)`
	_, err := db.Exec(ctx, sqlOrderDetail, orderId, seatId)
	if err != nil {
		log.Println("InsertOrderDetail Error:", err.Error())
		return err
	}
	return nil
}

func (o OrderRepository) GetSeatsByScheduleID(ctx context.Context, db DBTX, scheduleId int) ([]model.Seat, error) {
	// ... (sqlStr omitted)
	sqlStr := `
		SELECT 
			se.id AS seat_id,
			se.row_letter,
			se.seat_number,
			se.seat_type,
			CASE 
				WHEN EXISTS (
					SELECT 1 
					FROM orders o 
					JOIN order_details od ON od.order_id = o.id 
					WHERE o.schedule_id = sch.id 
						AND o.payment_status = 'paid' 
						AND od.seat_id = se.id
				) THEN 'sold'
				ELSE 'available'
			END AS status
		FROM schedules sch
		INNER JOIN cinemas c ON sch.cinema_id = c.id
		INNER JOIN seats se ON se.cinema_id = c.id
		WHERE sch.id = $1
		ORDER BY se.row_letter, se.seat_number;
	`

	rows, err := db.Query(ctx, sqlStr, scheduleId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []model.Seat
	for rows.Next() {
		var s model.Seat
		err := rows.Scan(&s.SeatId, &s.RowLetter, &s.SeatNumber, &s.SeatType, &s.Status)
		if err != nil {
			return nil, err
		}
		seats = append(seats, s)
	}
	return seats, nil
}

func (o OrderRepository) GetPriceFromSchedule(ctx context.Context, db DBTX, scheduleId int) (int, error) {
	sqlStr := "SELECT price FROM schedules WHERE id = $1"
	var price int
	err := db.QueryRow(ctx, sqlStr, scheduleId).Scan(&price)
	return price, err
}

func (o OrderRepository) UpdatePaymentStatus(ctx context.Context, db DBTX, orderId int, status string) error {
	sqlStr := "UPDATE orders SET payment_status = $1 WHERE id = $2"
	_, err := db.Exec(ctx, sqlStr, status, orderId)
	if err != nil {
		log.Println("UpdatePaymentStatus Error:", err.Error())
		return err
	}
	return nil
}
