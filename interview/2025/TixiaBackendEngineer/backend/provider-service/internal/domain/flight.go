package domain

const (
	StreamFlightSearchRequested = "flight.search.requested"
	StreamFlightSearchResults   = "flight.search.results"
)

type Flight struct {
	ID            string  `json:"id"`
	Airline       string  `json:"airline"`
	FlightNumber  string  `json:"flight_number"`
	From          string  `json:"from"`
	To            string  `json:"to"`
	DepartureTime string  `json:"departure_time"`
	ArrivalTime   string  `json:"arrival_time"`
	Price         float64 `json:"price"`
	Currency      string  `json:"currency"`
	Available     bool    `json:"available"`
}

type FlightSearchRequest struct {
	SearchID string `json:"search_id"`
	From     string `json:"from"`
	To       string `json:"to"`
	Date     string `json:"date"`
}
