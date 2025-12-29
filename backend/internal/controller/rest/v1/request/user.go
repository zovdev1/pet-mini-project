package request

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email string `join:"email"`
}
