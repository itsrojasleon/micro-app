package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rojasleon/reserve-micro/auth/models"
)

func TestSignupFlow(t *testing.T) {
	models.ConnectToDatabase("test-dev.db")

	router := SetupRouter()

	w := httptest.NewRecorder()
	body := bytes.NewBuffer([]byte(`{"email": "test@test.com", "password": "password"}`))

	var users []models.User
	models.DB.Find(&users)
	assert.Equal(t, 0, len(users), "Make sure DB is empty at the beggining")

	req, _ := http.NewRequest("POST", "/auth/signup", body)
	router.ServeHTTP(w, req)

	models.DB.Find(&users)
	assert.Equal(t, 1, len(users), "User was created with providing valid credentials")

	req, _ = http.NewRequest("POST", "/auth/signup", body)
	router.ServeHTTP(w, req)

	models.DB.Find(&users)
	assert.Equal(t, 1, len(users), "Cannot create another user with existing credentials")

	assert.NotEqual(t, "", w.Header().Get("Authorization"), "Creates a JWT")

	assert.Equal(t, http.StatusCreated, w.Code, "Returns a successful status code")

	// Clean up resources after tests ran
	models.DB.Migrator().DropTable(&models.User{})
}
