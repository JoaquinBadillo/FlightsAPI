/* seed.sql
 
Creates dummy data to start the database in a single transaction.
Requires the tables in the schema.sql file to be created first.

Joaquin Badillo
2024-04-14
*/

BEGIN;

INSERT INTO airports(icao, iata, name, state, country) VALUES ('KATL', 'ATL', 'Hartsfield-Jackson Atlanta International Airport', 'Atlanta', 'United States');
INSERT INTO airports(icao, iata, name, state, country) VALUES ('KLAX', 'LAX', 'Los Angeles International Airport', 'Los Angeles', 'United States');
INSERT INTO airports(icao, iata, name, state, country) VALUES ('MMMX', 'MEX', 'Licenciado Benito Juarez International Airport', 'Mexico City', 'Mexico');
INSERT INTO airports(icao, iata, name, state, country) VALUES ('KDEN', 'DEN', 'Denver International Airport', 'Denver', 'United States');

INSERT INTO flights(departure_time, arrival_time, origin_airport_id, destination_airport_id, status) VALUES ('2024-06-30 08:00:00', '2024-06-30 10:00:00', 'KATL', 'KLAX', 'AVAILABLE');
INSERT INTO flights(departure_time, arrival_time, origin_airport_id, destination_airport_id, status) VALUES ('2024-06-30 08:00:00', '2024-06-30 11:00:00', 'MMMX', 'KLAX', 'AVAILABLE');
INSERT INTO flights(departure_time, arrival_time, origin_airport_id, destination_airport_id, status) VALUES ('2021-03-24 08:00:00', '2021-03-24 10:00:00', 'KDEN', 'KATL', 'DEPARTED');
INSERT INTO flights(departure_time, arrival_time, origin_airport_id, destination_airport_id, status) VALUES ('2023-08-1 08:00:00', '2023-08-1 10:00:00', 'KDEN', 'KATL', 'CANCELLED');

-- Fully booked first flight
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (1, '1A', 'FIRST', 1000.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (1, '1B', 'FIRST', 1000.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (1, '2A', 'BUSINESS', 800.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (1, '2B', 'BUSINESS', 800.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (1, '3A', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (1, '3B', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (1, '4A', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (1, '4B', 'ECONOMY', 500.00);

INSERT INTO orders(email, first_name, last_name, payment_status) VALUES ('john.doe@mail.com', 'John', 'Doe', 'SUCCESSFUL');
UPDATE seats SET order_id = 1 WHERE flight_id = 1 AND seat_number = '1A';

INSERT INTO orders(email, first_name, last_name, payment_status) VALUES ('jane.doe@mail.com', 'Jane', 'Doe', 'SUCCESSFUL');
UPDATE seats SET order_id = 2 WHERE flight_id = 1 AND seat_number = '1B';

INSERT INTO orders(email, first_name, last_name, payment_status) VALUES ('fred@mysteryinc.com', 'Fred', 'Jones', 'SUCCESSFUL');
UPDATE seats SET order_id = 3 WHERE flight_id = 1 AND seat_number = '2A';

INSERT INTO orders(email, first_name, last_name, payment_status) VALUES ('velma@mysteryinc.com', 'Velma', 'Dinkley', 'SUCCESSFUL');
UPDATE seats SET order_id = 4 WHERE flight_id = 1 AND seat_number = '2B';

INSERT INTO orders(email, first_name, last_name, payment_status) VALUES ('shaggy@mysteryinc.com', 'Shaggy', 'Rogers', 'SUCCESSFUL');
UPDATE seats SET order_id = 5 WHERE flight_id = 1 AND seat_number = '3A';

INSERT INTO orders(email, first_name, last_name, payment_status) VALUES ('daphne@mysteryinc.com', 'Daphne', 'Blake', 'SUCCESSFUL');
UPDATE seats SET order_id = 6 WHERE flight_id = 1 AND seat_number = '3B';

INSERT INTO orders(email, first_name, last_name, payment_status) VALUES ('scoob@mysteryinc.com', 'Scooby', 'Doo', 'SUCCESSFUL');
UPDATE seats SET order_id = 7 WHERE flight_id = 1 AND seat_number = '4A';

INSERT INTO orders(email, first_name, last_name, payment_status) VALUES ('random@mail.com', 'Random', 'Person', 'SUCCESSFUL');
UPDATE seats SET order_id = 8 WHERE flight_id = 1 AND seat_number = '4B';

-- Second flight is available :)
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '1A', 'FIRST', 1000.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '1B', 'FIRST', 1000.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '2A', 'BUSINESS', 800.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '2B', 'BUSINESS', 800.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '3A', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '3B', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '4A', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '4B', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '5A', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (2, '5B', 'ECONOMY', 500.00);

-- Third flight already departed (Won't bother writing orders to seed db)
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '1A', 'FIRST', 1000.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '1B', 'FIRST', 1000.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '2A', 'BUSINESS', 800.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '2B', 'BUSINESS', 800.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '3A', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '3B', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '4A', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '4B', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '5A', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (3, '5B', 'ECONOMY', 500.00);

-- Fourth flight was cancelled
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (4, '1A', 'FIRST', 1000.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (4, '1B', 'FIRST', 1000.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (4, '2A', 'BUSINESS', 800.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (4, '2B', 'BUSINESS', 800.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (4, '3A', 'ECONOMY', 500.00);
INSERT INTO seats(flight_id, seat_number, class, price) VALUES (4, '3B', 'ECONOMY', 500.00);

COMMIT;