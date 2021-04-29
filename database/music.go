package database

import (
	"fmt"
	"musicapi/models"
	"strings"
)

//GetMusicsFiltered a
func GetMusicsFiltered(request models.MusicRequest) (models.MusicResponse, error) {
	musicResponse := models.MusicResponse{
		Musics: []models.Music{},
	}
	err := postgres.Select(&musicResponse.Musics, "SELECT * FROM musics "+
		"WHERE ((LOWER(name) LIKE $1) OR (LOWER(artist) LIKE $1)) "+
		"AND ($2='' OR genre=$2) "+
		"AND ($3='' OR album=$3)",
		"%"+strings.ToLower(request.Search)+"%", request.Genre, request.Album)

	return musicResponse, err
}

//GetMusic a
func GetMusic(id int) (models.Music, error) {
	var music models.Music
	err := postgres.Get(&music, "SELECT * FROM musics WHERE id=$1", id)
	return music, err
}

//InsertMusic a
func InsertMusic(music models.Music) error {
	_, err := postgres.Exec("INSERT INTO musics(name, artist, url, album, length, genre) VALUES($1, $2, $3, $4, $5, $6)",
		music.MusicName, music.ArtistName, music.URL, music.AlbumName, music.Length, music.Genre)
	fmt.Println(fmt.Sprint(err))
	return err
}

//DeleteMusic a
func DeleteMusic(id int) error {
	_, err := postgres.Exec("DELETE FROM musics WHERE id=$1", id)
	return err
}

/*func RegisterUser(user models.User) error {
	_, err := Postgres.Exec("INSERT INTO users(name, surname, email, password, registered_at) VALUES($1, $2, $3, $4, $5)", user.Name, user.Surname, user.Email, user.Password, user.RegisteredAt)
	return err
}

//GetUserByID ad
func GetUserByID(id int) models.User {
	user := models.User{}
	err := Postgres.Get(&user, "SELECT * FROM users WHERE ID=$1", id)
	if err != nil {
		panic(err)
	}
	return user
}
*/
