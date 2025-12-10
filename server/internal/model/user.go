package model

type User struct {
	ID       int64  `json:"id" db: "id"`
	Username string `json:"username" db: "username"`
	Email    string `json:"email" db: "email"`
	Password string `json:"password" db: "password"`
}

type RegisterReq struct {
	Username string `json:"username" db: "username"`
	Email    string `json:"email" db: "email"`
	Password string `json:"password" db: "password"`
}

type RegisterRes struct {
	ID       string `json:"id" db: "id"`
	Username string `json:"username" db: "username"`
	Email    string `json:"email" db: "email"`
}
