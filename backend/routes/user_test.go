package routes

import (
	"backend/database"
	"backend/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupApp() *fiber.App {
	app := fiber.New()
	app.Post("/api/register", Register)
	app.Post("/api/login", Login)
	app.Get("/api/users", GetUsers)
	return app
}

func TestRegister(t *testing.T) {
	// Setup
	database.ConnectDb()
	app := setupApp()

	// Test data
	user := models.User{
		Email:     "test@example.com",
		Password:  "password",
		FirstName: "John",
		LastName:  "Doe",
	}

	jsonUser, _ := json.Marshal(user)

	// Create request
	req := httptest.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var responseUser User
	json.NewDecoder(resp.Body).Decode(&responseUser)

	assert.Equal(t, user.Email, responseUser.Email)
	assert.NotEmpty(t, responseUser.ID)
	assert.NotEmpty(t, responseUser.Password)
	assert.NotEmpty(t, responseUser.FirstName)
	assert.NotEmpty(t, responseUser.LastName)

	// Cleanup
	cleanupDb()
}

func TestGetUsers(t *testing.T) {
	// Setup
	database.ConnectDb()
	app := setupApp()

	// Insert test data
	user := models.User{
		Email:     "test@example.com",
		Password:  "password",
		FirstName: "Jane",
		LastName:  "Doe",
	}
	database.Database.Db.Create(&user)

	// Create request
	req := httptest.NewRequest("GET", "/api/users", nil)

	// Perform request
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseUsers []User
	json.NewDecoder(resp.Body).Decode(&responseUsers)

	assert.Equal(t, user.Email, responseUsers[0].Email)
	assert.NotEmpty(t, responseUsers[0].ID)
	assert.NotEmpty(t, responseUsers[0].FirstName)
	assert.NotEmpty(t, responseUsers[0].LastName)

	// Cleanup
	cleanupDb()
}

func TestLogin(t *testing.T) {
	// Setup
	database.ConnectDb()
	app := setupApp()

	// Insert test data
	user := models.User{
		Email:     "test@example.com",
		Password:  "password",
		FirstName: "Toto",
		LastName:  "Doe",
	}
	database.Database.Db.Create(&user)

	// Test data
	loginInput := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    user.Email,
		Password: "password",
	}

	jsonLoginInput, _ := json.Marshal(loginInput)

	// Create request
	req := httptest.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonLoginInput))
	req.Header.Set("Content-Type", "application/json")

	// Perform request
	resp, err := app.Test(req)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var responseUser User
	json.NewDecoder(resp.Body).Decode(&responseUser)

	assert.Equal(t, user.Email, responseUser.Email)
	assert.NotEmpty(t, responseUser.ID)

	// Cleanup
	cleanupDb()
}

func cleanupDb() {
	database.Database.Db.Exec("DELETE FROM users")
}
