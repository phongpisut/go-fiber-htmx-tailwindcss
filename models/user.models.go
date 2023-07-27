package models

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Tel     string `json:"tel"`
	Option  uint64 `json:"option"`
	Message string `json:"message"`
}
