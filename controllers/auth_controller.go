package controllers

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Siginup(ctx *gin.Context)
}

type authController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) IAuthController {
	return &authController{service: service}
}

func (c *authController) Siginup(ctx *gin.Context) {
	var input dto.SiginupInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.Siginup(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	ctx.Status(http.StatusCreated)
}