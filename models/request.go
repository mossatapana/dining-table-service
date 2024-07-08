package models

type InitialRequest struct {
	NoOfTables int `json:"no_of_tables"`
}

type ReserveRequest struct {
	NoOfCustomers int `json:"no_of_customers"`
}

type CancelRequest struct {
	BookingID string `json:"booking_id"`
}
