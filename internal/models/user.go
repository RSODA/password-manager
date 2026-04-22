package models

type Service struct {
	ServiceName string `json:"service_name"`
	User        User   `json:"user"`
}

type User struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}

type Response struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
