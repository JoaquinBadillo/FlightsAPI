/* Models

Defines structs to create json reponses

Joaquin Badillo
2024-04-15
*/

package models

type Airport struct {
	ICAO    string `json:"icao,omitempty"`
	IATA    string `json:"iata,omitempty"`
	Name    string `json:"name,omitempty"`
	State   string `json:"state,omitempty"`
	Country string `json:"country,omitempty"`
}

type Flight struct {
	ID            int      `json:"id,omitempty"`
	Origin        *Airport `json:"origin,omitempty"`
	Destination   *Airport `json:"destination,omitempty"`
	ArrivalTime   *string  `json:"arrival,omitempty"`
	DepartureTime *string  `json:"departure,omitempty"`
}

type Seat struct {
	Flight *Flight `json:"flight,omitempty"`
	Number string  `json:"number,omitempty"`
	Class  string  `json:"class,omitempty"`
	Price  float64 `json:"price,omitempty"`
}

type Order struct {
	ID            int     `json:"id,omitempty"`
	Email         *string `json:"email,omitempty"`
	FirstName     *string `json:"first_name,omitempty"`
	LastName      *string `json:"last_name,omitempty"`
	Seat          *Seat   `json:"seat,omitempty"`
	PaymentStatus *string `json:"payment_status,omitempty"`
}
