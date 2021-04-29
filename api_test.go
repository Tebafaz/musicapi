package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"musicapi/handlers"
	"musicapi/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestApi(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)
	handlers.MapRoutes(r)

	testMethodForOk(r, t, "/", nil, http.MethodGet, nil)
	testMethodForOk(r, t, "/music/1", nil, http.MethodGet, nil)
	testMethodForOk(r, t, "/artist/1", nil, http.MethodGet, nil)
	credentials := []byte(`{"login":"testuser","password":"testuser"}`)
	newMusic := []byte(`{"album":"Alchemy 2 (Liquicity Presents)",
		"artist":"Dan Dakota",
		"genre":"Dance/Electronic",
		"legth":"04:27",
		"name":"Try not to worry",
		"url":"https://www.youtube.com/watch?v=rh7eA3oH5Ig"}`)
	newArtist := []byte(`{
	"birth_date": "string",
	"city": "Calgary",
	"fullname": "Dan Dakota",
	"nickname": "Dan Dakota"
	}`)
	testMethodForOk(r, t, "/user/signup", credentials, http.MethodPost, nil)
	testMethodForOk(r, t, "/user/login", credentials, http.MethodPost, ctx)
	testMethodForOk(r, t, "/music", newMusic, http.MethodPost, ctx)
	testMethodForOk(r, t, "/artist", newArtist, http.MethodPost, ctx)

	testMethodForOk(r, t, "/music", nil, http.MethodGet, ctx)
	testMethodForOk(r, t, "/artist", nil, http.MethodGet, ctx)

	testMethodForOk(r, t, "/music?search=Try not to worry&genre=Dance/Electronic&album=Alchemy 2 (Liquicity Presents)", nil, http.MethodGet, ctx)

	artistID, exist := ctx.Get("artistid")
	if !exist {
		t.Fatalf("couldn't find id of test artist")
	}
	musicID, exist := ctx.Get("musicid")
	if !exist {
		t.Fatalf("couldn't find id of test music")
	}
	testMethodForOk(r, t, fmt.Sprintf("/music/%d", musicID), nil, http.MethodDelete, ctx)
	testMethodForOk(r, t, fmt.Sprintf("/artist/%d", artistID), nil, http.MethodDelete, ctx)

	testMethodForOk(r, t, "/user", nil, http.MethodDelete, ctx)
}

func testMethodForOk(r *gin.Engine, t *testing.T, path string, data []byte, method string, ctx *gin.Context) {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("couldn't create request: %v\n", err)
	}
	if ctx != nil {
		valueInterface, exist := ctx.Get("token")

		if value, ok := valueInterface.(string); exist && ok {
			req.Header.Add("Authorization", value)
		}
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if path == "/user/login" {
		var token models.TokenResponse
		json.Unmarshal(w.Body.Bytes(), &token)
		tokenString := "Bearer " + token.Token
		ctx.Set("token", tokenString)
	}
	if path == "/music" {
		var musics models.MusicResponse
		json.Unmarshal(w.Body.Bytes(), &musics)
		for _, value := range musics.Musics {
			if value.MusicName == "Try not to worry" {
				ctx.Set("musicid", value.ID)
			}
		}
	}
	if path == "/artist" {
		var artists models.ArtistResponse
		json.Unmarshal(w.Body.Bytes(), &artists)
		for _, value := range artists.Artist {
			if value.Nickname == "Dan Dakota" {
				ctx.Set("artistid", value.ID)
			}
		}
	}
	if w.Code == http.StatusOK {
		fmt.Printf("\npath: %s\tmethod: %s\tdata: %s\n", path, method, w.Body.String())
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
