package controller

import (
	"net/http"

	"github.com/eulbyvan/enigma-university/middleware"
	"github.com/eulbyvan/enigma-university/model"
	"github.com/eulbyvan/enigma-university/model/dto/res"
	"github.com/eulbyvan/enigma-university/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
	rg          *gin.RouterGroup
}

func NewUserController(userUseCase usecase.UserUseCase, rg *gin.RouterGroup) *UserController {
	return &UserController{userUseCase: userUseCase}
}

func (c *UserController) FindById(ctx *gin.Context) {
	userID := ctx.Param("id")

	var res res.CommonResponse

	user, err := c.userUseCase.FindById(userID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	res.Code = http.StatusOK
	res.Status = "Success"
	res.Message = "Retrieved data successfully"
	res.Data = user

	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) FindAll(ctx *gin.Context) {

	var res res.CommonResponse

	users, err := c.userUseCase.FindAll()

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	res.Code = http.StatusOK
	res.Status = "Success"
	res.Message = "Retrieved data users successfully"
	res.Data = users

	ctx.JSON(http.StatusOK, res)
}

func (c *UserController) Post(ctx *gin.Context) {

	var user model.User

	ctx.BindJSON(&user)

	var res res.CommonResponse

	err := c.userUseCase.Post(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res.Code = http.StatusCreated
	res.Status = "Success"
	res.Message = "Create data successfully"
	// res.Data = user

	ctx.JSON(res.Code, res)
}

func (c *UserController) Delete(ctx *gin.Context) {

	id := ctx.Param("id")

	// ctx.BindJSON(&user)

	// fmt.Println("ini user >", user)

	var res res.CommonResponse

	err := c.userUseCase.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res.Code = http.StatusOK
	res.Status = "Success"
	res.Message = "Delete data user successfully"
	// res.Data = user

	ctx.JSON(res.Code, res)
}

func (c *UserController) Update(ctx *gin.Context) {

	var user model.User

	id := ctx.Param("id")

	ctx.BindJSON(&user)

	var res res.CommonResponse

	err := c.userUseCase.Update(id, user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res.Code = http.StatusOK
	res.Status = "Success"
	res.Message = "Delete data user successfully"
	// res.Data = user

	ctx.JSON(res.Code, res)
}

func (c *UserController) Login(ctx *gin.Context) {

	var user model.User

	// id := ctx.Param("id")

	ctx.BindJSON(&user)

	var res res.CommonResponse

	token, err := c.userUseCase.Login(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	res.Code = http.StatusOK
	res.Status = "Success"
	res.Message = "User Login Successfully"
	// res.Data = user

	ctx.JSON(res.Code, res)
}

func (c *UserController) Logout(ctx *gin.Context) {

	token, errToken := ctx.Cookie("Authorization")

	if errToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not logged in yet!!"})
		return
	}

	var res res.CommonResponse

	err := c.userUseCase.Logout(token)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)

	res.Code = http.StatusOK
	res.Status = "Success"
	res.Message = "User Logout Successfully"
	// res.Data = user

	ctx.JSON(res.Code, res)
}

func (c *UserController) Route() { // ERROR DI SEKITAR SINI
	c.rg.POST("/login", c.Login)
	c.rg.GET("/logout", c.Logout)
	// c.rg.GET("/users", userCtrl.FindById)
	users := c.rg.Group("/users")
	{
		users.Use(middleware.AuthorizeMiddleware())
		users.GET("/:id", c.FindById)
		users.GET("", c.FindAll)
		users.POST("", c.Post)
		users.DELETE("/:id", c.Delete)
		users.PUT("/:id", c.Update)
	}
}
