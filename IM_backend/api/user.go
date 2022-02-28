package api

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response_login struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
