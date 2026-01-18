package repository

import (
	"context"
	"log"
	"time"

	"github.com/Albaihaqi354/Tickitz-BE/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrdersRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o OrderRepository) GetSchedules(ctx context.Context, movieId int, showDate *string, city *string) ([]model.GetSchedules, error) {
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

	rows, err := o.db.Query(ctx, sqlStr, movieId, showDate, city)
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

func (o OrderRepository) CreateOrder(ctx context.Context, order model.Order, seats []int) (int, string, time.Time, error) {
	tx, err := o.db.Begin(ctx)
	if err != nil {
		return 0, "", time.Time{}, err
	}
	defer tx.Rollback(ctx)

	sqlOrder := `
		INSERT INTO orders (user_id, schedule_id, total_price, payment_status, booking_code)
		VALUES ($1, $2, $3, $4, md5(random()::text))
		RETURNING id, booking_code, created_at`

	var id int
	var bookingCode string
	var createdAt time.Time

	err = tx.QueryRow(ctx, sqlOrder, order.UserId, order.ScheduleId, order.TotalPrice, order.PaymentStatus).Scan(&id, &bookingCode, &createdAt)
	if err != nil {
		log.Println("CreateOrder Error:", err.Error())
		return 0, "", time.Time{}, err
	}

	sqlOrderDetail := `INSERT INTO order_details (order_id, seat_id) VALUES ($1, $2)`
	for _, seatId := range seats {
		_, err := tx.Exec(ctx, sqlOrderDetail, id, seatId)
		if err != nil {
			log.Println("CreateOrderDetail Error:", err.Error())
			return 0, "", time.Time{}, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, "", time.Time{}, err
	}

	return id, bookingCode, createdAt, nil
}

func (o OrderRepository) GetSeatsByScheduleID(ctx context.Context, scheduleId int) ([]model.Seat, error) {
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

	rows, err := o.db.Query(ctx, sqlStr, scheduleId)
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

func (o OrderRepository) GetPriceFromSchedule(ctx context.Context, scheduleId int) (int, error) {
	sqlStr := "SELECT price FROM schedules WHERE id = $1"
	var price int
	err := o.db.QueryRow(ctx, sqlStr, scheduleId).Scan(&price)
	return price, err
}

func (o OrderRepository) UpdatePaymentStatus(ctx context.Context, orderId int, status string) error {
	sqlStr := "UPDATE orders SET payment_status = $1 WHERE id = $2"
	_, err := o.db.Exec(ctx, sqlStr, status, orderId)
	if err != nil {
		log.Println("UpdatePaymentStatus Error:", err.Error())
		return err
	}
	return nil
}
