package models

//Music a
type Music struct {
	ID         int    `json:"id" form:"id" db:"id"`
	ArtistName string `json:"artist" form:"artist" db:"artist"`
	MusicName  string `json:"name" form:"name" db:"name"`
	URL        string `json:"url" form:"url" db:"url"`
	AlbumName  string `json:"album" form:"album" db:"album"`
	Length     string `json:"legth" form:"length" db:"length"`
	Genre      string `json:"genre" form:"genre" db:"genre"`
}

//MusicResponse a
type MusicResponse struct {
	Musics []Music `json:"musics" form:"musics"`
}

//MusicRequest a
type MusicRequest struct {
	Search string `json:"search" form:"search"`
	Album  string `json:"album" form:"album"`
	Genre  string `json:"genre" form:"genre"`
}
