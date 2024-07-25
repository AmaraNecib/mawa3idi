package api

import "time"

type ReservationType struct {
	Name string `json:"name"`
}

type Role struct {
	Name string `json:"name"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DaysOfWork struct {
	Name          string `json:"name"`
	From          string `json:"from"`
	To            string `json:"to"`
	Limit         int32  `json:"limit"`
	CurrentNumber int32  `json:"curent_number"`
	Date          string `json:"date"`
}

type workdays struct {
	Saturday   bool   `json:"saturday"`
	Sunday     bool   `json:"sunday"`
	Monday     bool   `json:"monday"`
	Tuesday    bool   `json:"tuesday"`
	Wednesday  bool   `json:"wednesday"`
	Thursday   bool   `json:"thursday"`
	Friday     bool   `json:"friday"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	MaxClients int32  `json:"max_clients"`
}

type GetWorkDay struct {
	Name                  string `json:"name"`
	NumberOfReservistions int    `json:"number_of_reservations"`
	MaxClients            int    `json:"max_clients"`
	OpenToWork            bool   `json:"open_to_work"`
	Date                  string `json:"date"`
	StartTime             string `json:"start_time"`
	EndTime               string `json:"end_time"`
}

type Weekday struct {
	ID         int64     `json:"id"`
	ServiceID  int64     `json:"service_id"`
	Name       string    `json:"name"`
	DayID      int64     `json:"day_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	MaxClients int       `json:"max_clients"`
}

type MyReservation struct {
	ServiceID int64 `json:"service_id"`
	// time 2024-07-26
	Time string `json:"time"`
	// Time       `json:"time"`
}
