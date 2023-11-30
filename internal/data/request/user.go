package request

type UserCreate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
