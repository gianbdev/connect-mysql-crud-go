package controllers

import (
	"go-api-crud/auth"
	"go-api-crud/database"
	"go-api-crud/models"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	record := database.Instance.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(401, gin.H{"error": "User not found"})
		context.Abort()
		return
	}

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(401, gin.H{"error": "Invalid credentials"})
		context.Abort()
		return
	}

	accessToken, refreshToken, err := auth.GenerateTokenPair(user.Email, user.Username)
	if err != nil {
		context.JSON(500, gin.H{"error": "Error generating token"})
		context.Abort()
		return
	}

	context.JSON(200, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func RefreshToken(context *gin.Context) {
	var request RefreshTokenRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := auth.ValidateToken(request.RefreshToken)
	if err != nil {
		context.JSON(401, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	claims, ok := token.Claims.(*auth.JWTClaim)
	// Can theoraically never happen since we already validated the token
	// maybe ValidateToken should return the claims instead of the token
	if !ok {
		context.JSON(401, gin.H{"error": "Invalid token"})
		context.Abort()
		return
	}

	// Generate new access token with the same claims
	accessToken, refreshToken, err := auth.GenerateTokenPair(claims.Email, claims.Username)
	if err != nil {
		context.JSON(500, gin.H{"error": "Error generating token"})
		context.Abort()
		return
	}

	context.JSON(200, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}
