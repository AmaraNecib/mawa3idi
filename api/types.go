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
