/* Database Provider

Defines a Database Manager using a singleton pattern.
The Manager type has an interface that defines the expected queries.

Joaquin Badillo
2024-04-15
*/

package provider

import (
	"database/sql"
	"errors"
	"log"
	"os"

	models "github.com/JoaquinBadillo/FlightsAPI/db/models"
	_ "github.com/lib/pq"
)

type Manager interface {
	GetFlight(id int) (*models.Flight, error)
	GetAvailableFlights(limit, offset int) ([]*models.Flight, error)
	GetAvailableFlightsByLocation(state, country string, limit, offset int) ([]*models.Flight, error)
	GetAvailableSeats(flightID int) ([]*models.Seat, error)
	CreateOrder(order *models.Order) (*models.Order, error)
	Close()
}

type manager struct {
	db *sql.DB
}

var Mgr Manager

func Connect() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_STRING"))

	if err != nil {
		panic(err)
	}

	log.Println("💾 Connected to database!")

	Mgr = &manager{db: db}
}

func (m *manager) GetFlight(id int) (*models.Flight, error) {
	origin := &models.Airport{}
	destination := &models.Airport{}
	f := &models.Flight{
		Origin:      origin,
		Destination: destination,
	}

	query := `
		SELECT f.id, f.arrival_time, f.departure_time, 
		a.icao, a.iata, a.name, a.state, a.country,
		b.icao, b.iata, b.name, b.state, b.country FROM flights f
		INNER JOIN airports a ON f.origin_airport_id = a.icao
		INNER JOIN airports b ON f.destination_airport_id = b.icao
		WHERE f.id = $1
	`
	err := m.db.QueryRow(query, id).Scan(
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

func (m *manager) GetAvailableFlights(limit, offset int) ([]*models.Flight, error) {
	query := `
		SELECT flight_id, arrival_time, departure_time, departure_state, 
		departure_country, arrival_state, arrival_country FROM available_flights
		LIMIT $1 OFFSET $2
	`

	rows, err := m.db.Query(query, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	flights := []*models.Flight{}

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
			return nil, err
		}

		flights = append(flights, f)
	}

	return flights, nil
}

func (m *manager) GetAvailableFlightsByLocation(state, country string, limit, offset int) ([]*models.Flight, error) {
	query := `
		SELECT flight_id, arrival_time, departure_time, departure_state, 
		departure_country, arrival_state, arrival_country FROM available_flights
		WHERE departure_state = $1 OR arrival_state = $1 AND 
		departure_country = $2 OR arrival_country = $2
		LIMIT $3 OFFSET $4
	`

	rows, err := m.db.Query(query, state, country, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	flights := []*models.Flight{}

	for rows.Next() {
		f := &models.Flight{}

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
			return nil, err
		}

		flights = append(flights, f)
	}

	return flights, nil
}

func (m *manager) GetAvailableSeats(flightID int) ([]*models.Seat, error) {
	query := `
		SELECT seat_number, class, price
		FROM available_seats
		WHERE flight_id = $1
	`

	rows, err := m.db.Query(query, flightID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	seats := []*models.Seat{}

	for rows.Next() {
		s := &models.Seat{}

		err := rows.Scan(
			&s.Number,
			&s.Class,
			&s.Price,
		)

		if err != nil {
			return nil, err
		}

		seats = append(seats, s)
	}

	return seats, nil
}

func (m *manager) CreateOrder(order *models.Order) (*models.Order, error) {
	query := "SELECT o_order_id, o_price FROM create_order($1, $2, $3, $4, $5)"

	if order.Email == nil ||
		order.FirstName == nil ||
		order.LastName == nil ||
		order.Seat == nil ||
		order.Seat.Flight == nil ||
		order.Seat.Flight.ID < 1 ||
		order.Seat.Number == "" {
		return nil, errors.New("invalid order data")
	}

	err := m.db.QueryRow(
		query,
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
	m.db.Close()
	log.Println("🔒 Closed database connection")
}
