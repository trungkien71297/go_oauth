package user

type User struct {
	Id        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
