package handlers

import (
	"fmt"
	"musicapi/database"
	"musicapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddArtist godoc
// @Summary Add artist
// @Description Inserts artist's data to database
// @Tags Artists
// @Accept  json
// @Produce  json
// @Param data body models.Artist true "Inserts this data into database"
// @Success 200 {object} models.SuccessJSON
// @Failure 400,401,500 {object} models.ErrorJSON
// @Router /artist [post]
func AddArtist(c *gin.Context) {
	var artist models.Artist
	err := c.ShouldBindJSON(&artist)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Bad request",
		})
		return
	}
	err = database.InsertArtist(artist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "error on the server side, please, contact with administrator",
		})
		return
	}
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusUnauthorized, models.ErrorJSON{
			Error: "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessJSON{
		Success: fmt.Sprintf("artist inserted to database by %s!", username.(string)),
	})
}

// GetArtist godoc
// @Summary Get artist
// @Description Get particular artist by its id
// @Tags Artists
// @Produce  json
// @Param id path int true "Artist ID"
// @Success 200 {object} models.Artist
// @Failure 404,400 {object} models.ErrorJSON
// @Router /artist/{id} [get]
func GetArtist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "passed not integer value",
		})
		return
	}
	artist, err := database.GetArtist(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorJSON{
			Error: "there is no artist with such id",
		})
		return
	}
	c.JSON(http.StatusOK, artist)
}

// GetArtists godoc
// @Summary Get all artists
// @Tags Artists
// @Produce  json
// @Success 200 {object} models.ArtistResponse
// @Failure 500 {object} models.ErrorJSON
// @Router /artist [get]
func GetArtists(c *gin.Context) {
	artists, err := database.GetArtists()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "error on the server side, please, contact with administrator",
		})
		return
	}
	c.JSON(http.StatusOK, artists)
}

//DeleteArtist godoc
// @Summary Delete artist
// @Description Delete artist by it's id
// @Tags Artists
// @Produce  json
// @Param id path int true "Artist ID"
// @Success 200 {object} models.SuccessJSON
// @Failure 400,500 {object} models.ErrorJSON
// @Router /artist/{id} [delete]
func DeleteArtist(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "passed not integer value",
		})
		return
	}
	err = database.DeleteArtist(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "Couldn't delete from database",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessJSON{
		Success: "Artist successfully deleted",
	})
}
