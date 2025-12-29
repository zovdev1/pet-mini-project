package dto

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Email    string
	Password string
}

type LogInUser struct {
	Email string
}
