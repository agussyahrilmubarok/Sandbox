package rest

// Auth user request payload
type AuthUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// User credential response payload
type UserCredential struct {
	Token string      `json:"token"`
	User  UserProfile `json:"user"`
}
