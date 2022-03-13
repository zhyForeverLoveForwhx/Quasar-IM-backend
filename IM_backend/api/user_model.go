package api

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Request_register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response_login struct {
	AccessToken           string        `json:"access_token"`
	Username string `json:"username"`
}

type Request_login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
