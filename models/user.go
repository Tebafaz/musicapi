package models

//User a
type User struct {
	ID       int    `json:"-" form:"-" db:"id"`
	Username string `json:"login" form:"login" db:"name"`
	Password string `json:"password" form:"password" db:"password"`
}
