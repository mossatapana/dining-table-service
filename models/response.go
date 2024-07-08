package models

type CommonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type InitialResponse struct {
	CommonResponse
}

type ReserveData struct {
	BookingID       string `json:"booking_id"`
	BookedTables    int    `json:"booked_tables"`
	RemainingTables int    `json:"remaining_tables"`
}

type ReserveResponse struct {
	CommonResponse
	Data ReserveData `json:"data"`
}

type CancelData struct {
	FreedTables     int `json:"freed_tables"`
	RemainingTables int `json:"remaining_tables"`
}

type CancelResponse struct {
	CommonResponse
	Data CancelData `json:"data"`
}
