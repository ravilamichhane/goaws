package types

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
