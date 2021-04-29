package models

//Artist a
type Artist struct {
	ID        int    `json:"id" form:"id" db:"id"`
	Nickname  string `json:"nickname" form:"nickname" db:"nickname"`
	FullName  string `json:"fullname" form:"fullname" db:"fullname"`
	City      string `json:"city" form:"city" db:"city"`
	BirthDate string `json:"birth_date" form:"birth_date" db:"birth_date"`
}

//ArtistRequest a
type ArtistRequest struct {
	NickName string `json:"nickname" form:"nickname"`
}

//ArtistResponse a
type ArtistResponse struct {
	Artist []Artist `json:"artists" form:"artists"`
}
