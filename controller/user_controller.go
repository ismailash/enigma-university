package controller

import (
	"net/http"

	"github.com/eulbyvan/enigma-university/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(userUseCase usecase.UserUseCase) *UserController {
	return &UserController{userUseCase: userUseCase}
}

func (c *UserController) FindById(ctx *gin.Context) {
	userID := ctx.Query("id")

	user, err := c.userUseCase.FindById(userID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, user)
}
