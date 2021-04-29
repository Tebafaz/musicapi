package database

import (
	"musicapi/models"
)

//GetUser a
func GetUser(request models.TokenRequest) (models.User, error) {
	var user models.User
	err := postgres.Get(&user, "SELECT * FROM users WHERE name=$1", request.Login)
	return user, err
}

//InsertUser a
func InsertUser(user models.User) error {
	_, err := postgres.Exec("INSERT INTO users(name, password) VALUES($1, $2)", user.Username, user.Password)
	return err
}

//DeleteUser a
func DeleteUser(username string) error {
	_, err := postgres.Exec("DELETE FROM users WHERE name=$1", username)
	return err
}
