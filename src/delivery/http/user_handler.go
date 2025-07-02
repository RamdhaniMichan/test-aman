package http

import (
	"net/http"
	"os"
	"test-aman/src/domain"
	"test-aman/src/lib"
	"test-aman/src/usecase"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		lib.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	createdUser, err := h.userUsecase.Register(&user)
	if err != nil {
		lib.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lib.SuccessResponse(c, http.StatusCreated, "User registered successfully", createdUser)
}

func (h *UserHandler) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		lib.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userUsecase.Login(credentials.Email, credentials.Password)
	if err != nil {
		lib.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	lib.SuccessResponse(c, http.StatusOK, "Login successful", gin.H{
		"token": tokenString,
		"user":  user,
	})
}
