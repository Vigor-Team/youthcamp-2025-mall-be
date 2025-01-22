package types

type Token struct {
	Token string `json:"token"`
}

type UserInfo struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}
