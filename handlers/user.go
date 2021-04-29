package handlers

import (
	"musicapi/database"
	"musicapi/models"
	"musicapi/redis"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login
// @Description Gives the token to requested machine, if it provided valid login and password
// @Tags Users
// @Accept  json
// @Produce  json
// @Param data body models.TokenRequest true "Requires login and password in order to give token"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorJSON
// @Router /login [post]
func Login(c *gin.Context) {
	var request models.TokenRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Bad request",
		})
		return
	}
	if request.Login == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Login or password is empty",
		})
		return
	}
	user, err := database.GetUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Username and/or password incorrect",
		})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Username and/or password incorrect",
		})
		return
	}
	token := uuid.New().String()
	redis.AddToken(user, token, 60*60)
	c.JSON(http.StatusOK, models.TokenResponse{
		Token: token,
	})
}

// Signup godoc
// @Summary Sign up
// @Description Inserts user to database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param data body models.User true "Inserts this data into database"
// @Success 200 {object} models.SuccessJSON
// @Failure 400,500 {object} models.ErrorJSON
// @Router /signup [post]
func Signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Bad request",
		})
		return
	}
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Login or password is empty",
		})
		return
	}
	bytePass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorJSON{
			Error: "Bad request",
		})
		return
	}
	user.Password = string(bytePass)
	err = database.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "This username already exists",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessJSON{
		Success: "New user created!",
	})
}

//DeleteUser godoc
// @Summary Delete user
// @Description Delete user that holds token
// @Tags Users
// @Produce  json
// @Param id path int true "Artist ID"
// @Success 200 {object} models.SuccessJSON
// @Failure 401,500 {object} models.ErrorJSON
// @Router /user [delete]
func DeleteUser(c *gin.Context) {
	username, exist := c.Get("username")

	if !exist {
		c.JSON(http.StatusUnauthorized, models.ErrorJSON{
			Error: "You can't delete your account, you are not logged in",
		})
		return
	}
	err := database.DeleteUser(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorJSON{
			Error: "Couldn't delete from database",
		})
		return
	}
	c.JSON(http.StatusOK, models.SuccessJSON{
		Success: "User successfully deleted",
	})
}
