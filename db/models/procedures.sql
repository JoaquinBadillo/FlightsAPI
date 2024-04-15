/* procedures.sql

Defines databases procedures using PL/pgSQL for PostgreSQL.

  Create Order:
  Checks if the seat is available for the flight.
  Creates an order and assigns it to the seat.

Joaquin Badillo
2024-04-15
*/

CREATE OR REPLACE FUNCTION create_order(
  p_email VARCHAR(255),
  p_first_name VARCHAR(255),
  p_last_name VARCHAR(255),
  p_flight_id INTEGER,
  p_seat_number VARCHAR(10),
  p_payment_status payment_status_enum DEFAULT 'PENDING',
  OUT o_order_id INTEGER,
  OUT o_price DECIMAL(10, 2)
) AS $$
BEGIN

  SELECT price INTO o_price FROM seats
  WHERE flight_id = p_flight_id AND seat_number = p_seat_number
  AND order_id IS NULL;
  
  IF o_price IS NULL THEN
    RAISE EXCEPTION 'Seat % is not available for flight %', p_seat_number, p_flight_id;
  END IF;

  INSERT INTO orders(email, first_name, last_name, payment_status)
  VALUES(p_email, p_first_name, p_last_name, p_payment_status)
  RETURNING id INTO o_order_id;

  UPDATE seats SET order_id = o_order_id
  WHERE flight_id = p_flight_id AND seat_number = p_seat_number 
  AND order_id IS NULL;
END; $$ LANGUAGE plpgsql;

