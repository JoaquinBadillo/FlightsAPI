/* Models

Defines structs to create json reponses

Joaquin Badillo
2024-04-14
*/

package models

type Airport struct {
	ICAO    string `json:"icao,omitempty"`
	IATA    string `json:"iata,omitempty"`
	Name    string `json:"name,omitempty"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type Flight struct {
	ID            int     `json:"id,omitempty"`
	Origin        Airport `json:"origin"`
	Destination   Airport `json:"destination"`
	ArrivalTime   string  `json:"arrival"`
	DepartureTime string  `json:"departure"`
}

type Seat struct {
	Flight Flight  `json:"flight"`
	Number string  `json:"number"`
	Class  string  `json:"class"`
	Status string  `json:"status"`
	Price  float64 `json:"price"`
}
