package models

//TokenRequest a
type TokenRequest struct {
	Login    string `json:"login" form:"login"`
	Password string `json:"password" form:"password"`
}

//TokenResponse a
type TokenResponse struct {
	Token string `json:"token"`
}
