package api

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response_login struct {
	Username    string `json:"username"`
	AccessToken string `json:"token"`
}

type Request_login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
