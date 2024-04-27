/* Database Provider

Defines a Database Manager using a singleton pattern.
The Manager type has an interface that defines the expected queries.

Joaquin Badillo
2024-04-15
*/

package provider

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	models "github.com/JoaquinBadillo/FlightsAPI/db/models"
	_ "github.com/lib/pq"
)

type Manager interface {
	GetFlight(ctx context.Context, id int) (*models.Flight, error)
	GetAvailableFlights(ctx context.Context, limit, offset int) ([]*models.Flight, error)
	GetAvailableFlightsByLocation(ctx context.Context, state, country string, limit, offset int) ([]*models.Flight, error)
	GetAvailableSeats(ctx context.Context, flightID int) ([]*models.Seat, error)
	CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error)
	Close()
}

type manager struct {
	db                                *sql.DB
	getFlightStmt                     *sql.Stmt
	getAvailableFlightsStmt           *sql.Stmt
	getAvailableFlightsByLocationStmt *sql.Stmt
	getAvailableSeatsStmt             *sql.Stmt
	createOrderStmt                   *sql.Stmt
}

var Mgr Manager

func Connect() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_STRING"))

	if err != nil {
		panic(err)
	}

	log.Println("💾 Connected to database!")

	m := &manager{db: db}

	log.Println("🍳 Preparing statements...")

	m.getFlightStmt, err = db.Prepare(`
		SELECT f.id, f.arrival_time, f.departure_time,
		a.icao, a.iata, a.name, a.state, a.country,
		b.icao, b.iata, b.name, b.state, b.country FROM flights f
		INNER JOIN airports a ON f.origin_airport_id = a.icao
		INNER JOIN airports b ON f.destination_airport_id = b.icao
		WHERE f.id = $1
	`)

	if err != nil {
		panic(err)
	}

	m.getAvailableFlightsStmt, err = db.Prepare(`
		SELECT flight_id, arrival_time, departure_time, departure_state,
		departure_country, arrival_state, arrival_country FROM available_flights
		LIMIT $1 OFFSET $2
	`)

	if err != nil {
		panic(err)
	}

	m.getAvailableFlightsByLocationStmt, err = db.Prepare(`
		SELECT flight_id, arrival_time, departure_time, departure_state,
		departure_country, arrival_state, arrival_country FROM available_flights
		WHERE departure_state = $1 OR arrival_state = $1 AND
		departure_country = $2 OR arrival_country = $2
		LIMIT $3 OFFSET $4
	`)

	if err != nil {
		panic(err)
	}

	m.getAvailableSeatsStmt, err = db.Prepare(`
		SELECT seat_number, class, price
		FROM available_seats
		WHERE flight_id = $1
	`)

	if err != nil {
		panic(err)
	}

	m.createOrderStmt, err = db.Prepare(
		"SELECT o_order_id, o_price FROM create_order($1, $2, $3, $4, $5)",
	)

	if err != nil {
		panic(err)
	}

	Mgr = m
}

func (m *manager) GetFlight(ctx context.Context, id int) (*models.Flight, error) {
	origin := &models.Airport{}
	destination := &models.Airport{}
	f := &models.Flight{
		Origin:      origin,
		Destination: destination,
	}

	err := m.getFlightStmt.QueryRowContext(ctx, id).Scan(
		&f.ID,
		&f.ArrivalTime,
		&f.DepartureTime,
		&origin.ICAO,
		&origin.IATA,
		&origin.Name,
		&origin.State,
		&origin.Country,
		&destination.ICAO,
		&destination.IATA,
		&destination.Name,
		&destination.State,
		&destination.Country,
	)

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (m *manager) GetAvailableFlights(ctx context.Context, limit, offset int) ([]*models.Flight, error) {
	rows, err := m.getAvailableFlightsStmt.QueryContext(
		ctx,
		min(30, limit),
		offset,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()

	flights := []*models.Flight{}

	errChan := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		for rows.Next() {
			f := &models.Flight{
				Origin:      &models.Airport{},
				Destination: &models.Airport{},
			}

			if err := rows.Scan(
				&f.ID,
				&f.ArrivalTime,
				&f.DepartureTime,
				&f.Origin.State,
				&f.Origin.Country,
				&f.Destination.State,
				&f.Destination.Country,
			); err != nil {
				errChan <- err
				return
			}

			flights = append(flights, f)
		}
		done <- struct{}{}
	}()

	select {
	case <-done:
		return flights, nil
	case <-ctx.Done():
		log.Printf("❌ %v", ctx.Err())
		return nil, ctx.Err()
	case err := <-errChan:
		return nil, err
	}
}

func (m *manager) GetAvailableFlightsByLocation(ctx context.Context, state, country string, limit, offset int) ([]*models.Flight, error) {
	rows, err := m.getAvailableFlightsByLocationStmt.QueryContext(
		ctx,
		state,
		country,
		min(30, limit),
		offset,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	flights := []*models.Flight{}

	errChan := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		for rows.Next() {
			f := &models.Flight{
				Origin:      &models.Airport{},
				Destination: &models.Airport{},
			}

			err := rows.Scan(
				&f.ID,
				&f.ArrivalTime,
				&f.DepartureTime,
				&f.Origin.State,
				&f.Origin.Country,
				&f.Destination.State,
				&f.Destination.Country,
			)

			if err != nil {
				errChan <- err
				return
			}

			flights = append(flights, f)
		}

		done <- struct{}{}
	}()

	select {
	case <-done:
		return flights, nil
	case <-ctx.Done():
		log.Printf("❌ %v", ctx.Err())
		return nil, ctx.Err()
	case err := <-errChan:
		return nil, err
	}
}

func (m *manager) GetAvailableSeats(ctx context.Context, flightID int) ([]*models.Seat, error) {
	rows, err := m.getAvailableSeatsStmt.QueryContext(ctx, flightID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	seats := []*models.Seat{}

	errChan := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		for rows.Next() {
			s := &models.Seat{}

			err := rows.Scan(
				&s.Number,
				&s.Class,
				&s.Price,
			)

			if err != nil {
				errChan <- err
				return
			}

			seats = append(seats, s)
		}

		done <- struct{}{}
	}()

	select {
	case <-done:
		return seats, nil
	case <-ctx.Done():
		log.Printf("❌ %v", ctx.Err())
		return nil, ctx.Err()
	case err := <-errChan:
		return nil, err
	}
}

func (m *manager) CreateOrder(ctx context.Context, order *models.Order) (*models.Order, error) {
	if order.Email == nil ||
		order.FirstName == nil ||
		order.LastName == nil ||
		order.Seat == nil ||
		order.Seat.Flight == nil ||
		order.Seat.Flight.ID < 1 ||
		order.Seat.Number == "" {
		return nil, errors.New("invalid order data")
	}

	err := m.createOrderStmt.QueryRowContext(
		ctx,
		order.Email,
		order.FirstName,
		order.LastName,
		order.Seat.Flight.ID,
		order.Seat.Number,
	).Scan(
		&order.ID,
		&order.Seat.Price,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (m *manager) Close() {
	m.getFlightStmt.Close()
	m.getAvailableFlightsStmt.Close()
	m.getAvailableFlightsByLocationStmt.Close()
	m.getAvailableSeatsStmt.Close()
	m.createOrderStmt.Close()
	log.Println("🔒 Closed prepared statements")

	m.db.Close()
	log.Println("🔒 Closed database connection")
}
