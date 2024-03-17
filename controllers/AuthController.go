package controllers

import (
	"go-gin-rest-boilerplate/middlewares"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

// TokenResponse represents the structure of the token response.
type TokenResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiIxMjMiLCJ1c2VybmFtZSI6IkJlYXN0Iiwicm9sZXMiOlsiYWRtaW4iLCJwcm8tdXNlciJdLCJleHAiOjE2NDg0MzUzMDAsImlzcyI6ImdvLWJlYXN0In0.Z6q2hmZYrRzvR6VZQab3t1Q-X3chVJjXHVtjvOVq1K0"`
}

// ErrorResponse represents the structure of the error response.
type ErrorResponse struct {
	Message string `json:"message" example:"An error occurred"`
}

// @Summary User login
// @Description Logs in a user and returns an access token
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} ErrorResponse
// @Router /login [post]
func (auth AuthController) HandleLogin(c *gin.Context) {

	userId := "123"
	username := "Beast"
	roles := []string{"admin", "pro-user"}

	// do user auth here

	//issue token
	token, err := middlewares.GenerateToken([]byte(middlewares.SigningKey), userId, username, roles)
	if err != nil {
		c.JSON(400, ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(200, TokenResponse{Token: token})
}
