/* seed.sql
 
Creates dummy data to start the database in a single transaction.
Requires the tables in the schema.sql file to be created first.

Joaquin Badillo
2024-04-14
*/

BEGIN;
SET session_replication_role = 'replica';

INSERT INTO airports(icao, iata, name, state, country) VALUES 
('KATL', 'ATL', 'Hartsfield-Jackson Atlanta International Airport', 'Atlanta', 'United States'),
('KLAX', 'LAX', 'Los Angeles International Airport', 'Los Angeles', 'United States'),
('MMMX', 'MEX', 'Licenciado Benito Juarez International Airport', 'Mexico City', 'Mexico'),
('KDEN', 'DEN', 'Denver International Airport', 'Denver', 'United States');

INSERT INTO flights(departure_time, arrival_time, origin_airport_id, destination_airport_id, status) VALUES 
('2024-06-30 08:00:00', '2024-06-30 10:00:00', 'KATL', 'KLAX', 'AVAILABLE'),
('2024-06-30 08:00:00', '2024-06-30 10:00:00', 'KATL', 'MMMX', 'AVAILABLE'),
('2021-03-24 08:00:00', '2021-03-24 10:00:00', 'KDEN', 'KATL', 'DEPARTED'),
('2023-08-1 08:00:00', '2023-08-1 10:00:00', 'KDEN', 'KATL', 'CANCELLED');

-- Fully booked first flight
INSERT INTO seats(flight_id, seat_number, class, price, order_id) VALUES
(1, '1A', 'FIRST', 1000.00, 1),
(1, '1B', 'FIRST', 1000.00, 2),
(1, '2A', 'BUSINESS', 800.00, 3),
(1, '2B', 'BUSINESS', 800.00, 4),
(1, '3A', 'ECONOMY', 500.00, 5),
(1, '3B', 'ECONOMY', 500.00, 6),
(1, '4A', 'ECONOMY', 500.00, 7),
(1, '4B', 'ECONOMY', 500.00, 8);

INSERT INTO orders(email, first_name, last_name, payment_status) VALUES 
('john.doe@mail.com', 'John', 'Doe', 'SUCCESSFUL'),
('fred@mysteryinc.com', 'Fred', 'Jones', 'SUCCESSFUL'),
('velma@mysteryinc.com', 'Velma', 'Dinkley', 'SUCCESSFUL'),
('shaggy@mysteryinc.com', 'Shaggy', 'Rogers', 'SUCCESSFUL'),
('daphne@mysteryinc.com', 'Daphne', 'Blake', 'SUCCESSFUL'),
('scoob@mysteryinc.com', 'Scooby', 'Doo', 'SUCCESSFUL'),
('random@mail.com', 'Random', 'Person', 'SUCCESSFUL');

-- Second flight is available :)
INSERT INTO seats(flight_id, seat_number, class, price) VALUES 
(2, '1A', 'FIRST', 1000.00),
(2, '1B', 'FIRST', 1000.00),
(2, '2A', 'BUSINESS', 800.00),
(2, '2B', 'BUSINESS', 800.00),
(2, '3A', 'ECONOMY', 500.00),
(2, '3B', 'ECONOMY', 500.00),
(2, '4A', 'ECONOMY', 500.00),
(2, '4B', 'ECONOMY', 500.00),
(2, '5A', 'ECONOMY', 500.00),
(2, '5B', 'ECONOMY', 500.00);

-- Third flight already departed (Won't bother writing orders to seed db)
INSERT INTO seats(flight_id, seat_number, class, price) VALUES 
(3, '1A', 'FIRST', 1000.00),
(3, '1B', 'FIRST', 1000.00),
(3, '2A', 'BUSINESS', 800.00),
(3, '2B', 'BUSINESS', 800.00),
(3, '3A', 'ECONOMY', 500.00),
(3, '3B', 'ECONOMY', 500.00),
(3, '4A', 'ECONOMY', 500.00),
(3, '4B', 'ECONOMY', 500.00),
(3, '5A', 'ECONOMY', 500.00),
(3, '5B', 'ECONOMY', 500.00);

-- Fourth flight was cancelled
INSERT INTO seats(flight_id, seat_number, class, price) VALUES
(4, '1A', 'FIRST', 1000.00),
(4, '1B', 'FIRST', 1000.00),
(4, '2A', 'BUSINESS', 800.00),
(4, '2B', 'BUSINESS', 800.00),
(4, '3A', 'ECONOMY', 500.00),
(4, '3B', 'ECONOMY', 500.00),
(4, '4A', 'ECONOMY', 500.00),
(4, '4B', 'ECONOMY', 500.00);

SET session_replication_role = 'origin';
COMMIT;