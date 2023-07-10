package api

import (
	"crypto/sha256"
	"fmt"
	"github.com/jonasdoesthings/plavatar"
	"github.com/jonasdoesthings/plavatar/plavatar-rest/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	avatarGenerator = &plavatar.Generator{}
	apiServer       = &Server{
		logger:          utils.InitLogger(),
		echoRouter:      echo.New(),
		avatarGenerator: avatarGenerator,
	}
)

// todo: do a snapshot test for each avatar type
func TestServer_HandleGetAvatar(t *testing.T) {
	minSize = 256
	maxSize = 1024

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := apiServer.echoRouter.NewContext(req, rec)
	c.SetPath("/smiley/:size/:name")
	c.SetParamNames("size", "name")
	c.SetParamValues("512", "6")

	var shaHasher = sha256.New()

	h := apiServer.HandleGetAvatar(avatarGenerator.Smiley)
	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "6", rec.Header().Get("Rng-Seed"))

		shaHasher.Write(rec.Body.Bytes())
		hash := fmt.Sprintf("%x", shaHasher.Sum(nil))
		assert.Equal(t, "45f2f471e4df40cf1fd2b424b1bb2275491e6b9198e3f70bbb9565519dcaa82b", hash)
	}
}
