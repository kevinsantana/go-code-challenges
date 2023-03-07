package models

// SignUp struct to describe register a new user.
type SignUp struct {
	Name     string `json:"name" validate:"required,name,lte=255"`
	CPF      string `json:"cpf" validate:"required,cpf,lte=255"`
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

// SignIn struct to describe login user.
type SignIn struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}
