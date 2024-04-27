/* schema.sql
 
Defines the database models for PostgreSQL in a single transaction.

Joaquin Badillo
2024-04-14
*/

BEGIN;

CREATE TABLE airports(
  icao VARCHAR(4) PRIMARY KEY,
  iata VARCHAR(3) UNIQUE,
  name VARCHAR(255) NOT NULL,
  country VARCHAR(255) NOT NULL,
  state VARCHAR(255) NOT NULL
);

CREATE INDEX idx_airports_iata ON airports(iata);
CREATE INDEX idx_airports_location ON airports(country, state);

CREATE TYPE flight_status_enum AS ENUM ('AVAILABLE', 'DEPARTED', 'CANCELLED');

CREATE TABLE flights(
  id SERIAL PRIMARY KEY,
  departure_time TIMESTAMP NOT NULL,
  arrival_time TIMESTAMP NOT NULL,
  origin_airport_id VARCHAR(4) REFERENCES airports(icao) ON DELETE CASCADE,
  destination_airport_id VARCHAR(4) REFERENCES airports(icao) ON DELETE CASCADE,
  status flight_status_enum NOT NULL
);

CREATE TYPE payment_status_enum AS ENUM ('PENDING', 'SUCCESSFUL');

CREATE TABLE orders(
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  payment_status payment_status_enum NOT NULL DEFAULT 'PENDING'
);

CREATE INDEX idx_orders_email ON orders(email);

CREATE TYPE seat_class_enum AS ENUM ('ECONOMY', 'BUSINESS', 'FIRST');

CREATE TABLE seats(
  flight_id INTEGER REFERENCES flights(id) ON DELETE CASCADE,
  seat_number VARCHAR(10) NOT NULL,
  class seat_class_enum NOT NULL DEFAULT 'ECONOMY',
  order_id INTEGER UNIQUE REFERENCES orders(id) ON DELETE SET NULL,
  price DECIMAL(10, 2) NOT NULL,
  PRIMARY KEY (flight_id, seat_number)
);

CREATE INDEX idx_seats_order_id ON seats(order_id);
CREATE INDEX idx_seats_flight_id ON seats(flight_id);

CREATE VIEW available_flights AS
SELECT DISTINCT f.id AS flight_id, f.departure_time, f.arrival_time,
departure.state AS departure_state, departure.country AS departure_country,
arrival.state AS arrival_state, arrival.country AS arrival_country
FROM flights f
JOIN airports departure ON f.origin_airport_id = departure.icao
JOIN airports arrival ON f.destination_airport_id = arrival.icao
JOIN seats s ON f.id = s.flight_id
WHERE s.order_id IS NULL
AND f.status = 'AVAILABLE';

CREATE VIEW booked_seats AS
SELECT f.id AS flight_id, f.departure_time, f.arrival_time,
departure.state AS departure_state, departure.country AS departure_country,
arrival.state AS arrival_state, arrival.country AS arrival_country,
s.seat_number, s.class, s.price
FROM flights f
JOIN airports departure ON f.origin_airport_id = departure.icao
JOIN airports arrival ON f.destination_airport_id = arrival.icao
JOIN seats s ON f.id = s.flight_id
WHERE s.order_id IS NOT NULL;

CREATE VIEW available_seats AS
SELECT f.id AS flight_id, f.departure_time, f.arrival_time,
departure.state AS departure_state, departure.country AS departure_country,
arrival.state AS arrival_state, arrival.country AS arrival_country,
s.seat_number, s.class, s.price
FROM flights f
JOIN airports departure ON f.origin_airport_id = departure.icao
JOIN airports arrival ON f.destination_airport_id = arrival.icao
JOIN seats s ON f.id = s.flight_id
WHERE s.order_id IS NULL;

COMMIT;