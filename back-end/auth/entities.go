package auth

type UserLogin struct {
	User     string `json: "user"`
	Password string `json: "password"`
}

type UserJTW struct {
	Token string `json: "user"`
}
