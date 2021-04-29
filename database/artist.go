package database

import (
	"musicapi/models"
)

//GetArtists a
func GetArtists() (models.ArtistResponse, error) {
	var artistResponse models.ArtistResponse
	err := postgres.Select(&artistResponse.Artist, "SELECT * FROM artists")
	return artistResponse, err
}

//GetArtist a
func GetArtist(id int) (models.Artist, error) {
	var artist models.Artist
	err := postgres.Get(&artist, "SELECT * FROM artists WHERE id=$1", id)
	return artist, err
}

//InsertArtist a
func InsertArtist(artist models.Artist) error {
	_, err := postgres.Exec("INSERT INTO artists(nickname, fullname, city, birth_date) VALUES($1, $2, $3, $4)",
		artist.Nickname, artist.FullName, artist.City, artist.BirthDate)
	return err
}

//DeleteArtist a
func DeleteArtist(id int) error {
	_, err := postgres.Exec("DELETE FROM artists WHERE id=$1", id)
	return err
}
