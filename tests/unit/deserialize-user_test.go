package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"tiny-site-backend/middleware"
	"tiny-site-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func getMockContextWithToken(token string) *gin.Context {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set the token in the request's cookies.
	c.Request.AddCookie(&http.Cookie{
		Name:  "token",
		Value: token,
	})

	return c
}

func TestDeserializeUser_DecodeToken(t *testing.T) {
	// Create a mock token string (for testing purposes).
	mockToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY5Mzk5NjUsImlhdCI6MTY5NjkzNjM2NSwibmJmIjoxNjk2OTM2MzY1LCJzdWIiOiI1YjhlODhjMi05NTVhLTRlMzMtYjIyOC02MzJkZTFmMzhmNDUifQ.lLFD1AWiDUUOtye1mnsD63X6gytuJxoYnfITZpV2DO8"

	// Create a mock UUID for the user ID.
	mockUserID := uuid.MustParse("5b8e88c2-955a-4e33-b228-632de1f38f45")

	// Create a mock User with the pointer to the UUID.
	mockUser := &models.User{
		ID:        &mockUserID,
		FirstName: "John",
		LastName:  "Doe",
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		// ... fill in other fields as needed ...
	}

	// Create a mock Gin context with the mock token.
	c := getMockContextWithToken(mockToken)

	// Set the mock user in the context.
	c.Set("user", mockUser)

	// Call the DeserializeUser middleware.
	middleware.DeserializeUser()(c)

	// Check if the "user" key exists in the context and matches the mock user.
	user, exists := c.Get("user")
	assert.True(t, exists, "Expected key 'user' to exist in the context")
	assert.Equal(t, mockUser, user.(*models.User), "Expected user to be set to the mock user")

	// Check that the response status code is 200.
	assert.Equal(t, http.StatusOK, c.Writer.Status())
}
