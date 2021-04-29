package handlers

import (
	"fmt"
	"musicapi/database"
	"musicapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMusics godoc
// @Summary Get multiple musics
// @Description Get all or filtered list of music. Filters are search, genre, album. Could be used partially [Not reqired]
// @Tags Musics
// @Produce  json
// @Param data query models.MusicRequest false "returns everything that matches all the passed parameters"
// @Success 200 {object} models.MusicResponse
// @Failure 400,500 {object} models.ErrorJSON
// @Router /music [get]
func GetMusics(c *gin.Context) {
	var request models.MusicRequest
	err := c.ShouldBind(&request)
	fmt.Println(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Bad request",
		})
		return
	}
	musics, err := database.GetMusicsFiltered(request)
	fmt.Printf("err: %v\n", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "server side error, please contact with administrator",
		})
		return
	}
	c.JSON(http.StatusOK, musics)
}

// GetMusic godoc
// @Summary Get music
// @Description Get particular music by its id
// @Tags Musics
// @Produce  json
// @Param id path int true "Music ID"
// @Success 200 {object} models.Music
// @Failure 404,400 {object} models.ErrorJSON
// @Router /music/{id} [get]
func GetMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "passed not integer value",
		})
		return
	}
	music, err := database.GetMusic(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorJSON{
			Error: "there is no artist with such id",
		})
		return
	}
	c.JSON(http.StatusOK, music)
}

// AddMusic godoc
// @Summary Insert music
// @Description Inserts music to database
// @Tags Musics
// @Accept  json
// @Produce  json
// @Param data body models.Music true "Inserts this data into database"
// @Success 200 {object} models.SuccessJSON
// @Failure 400,500 {object} models.ErrorJSON
// @Router /music [post]
func AddMusic(c *gin.Context) {
	var music models.Music
	err := c.ShouldBindJSON(&music)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Bad request",
		})
		return
	}
	err = database.InsertMusic(music)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "Server side error, contact with administrator",
		})
		return
	}
	username, exist := c.Get("username")
	if !exist {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessJSON{
		Success: fmt.Sprintf("music inserted to database by %s!", username.(string)),
	})
}

//DeleteMusic godoc
// @Summary Delete music
// @Description Delete music by it's id
// @Tags Musics
// @Produce  json
// @Param id path int true "Music ID"
// @Success 200 {object} models.SuccessJSON
// @Failure 400,500 {object} models.ErrorJSON
// @Router /music/{id} [delete]
func DeleteMusic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "passed not integer value",
		})
		return
	}
	err = database.DeleteMusic(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "Couldn't delete from database",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessJSON{
		Success: "Music successfully deleted",
	})
}
