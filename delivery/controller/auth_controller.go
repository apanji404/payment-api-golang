package controller

import (
	"mnc/delivery/middleware"
	"mnc/model"
	"mnc/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router  *gin.Engine
	usecase usecase.AuthUseCase
}

func (a *AuthController) loginHandler(c *gin.Context) {
	var payload model.Customer
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	token, err := a.usecase.Login(payload.Username, payload.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (a *AuthController) logoutHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := a.usecase.Logout(tokenString); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

func NewAuthController(usecase usecase.AuthUseCase, r *gin.Engine) *AuthController {
	controller := AuthController{
		router:  r,
		usecase: usecase,
	}
	rg := r.Group("/api")
	rg.POST("/login", controller.loginHandler)
	rg.POST("/logout", middleware.AuthMiddleware(), controller.logoutHandler)
	return &controller
}
