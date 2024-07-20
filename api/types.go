package api

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
