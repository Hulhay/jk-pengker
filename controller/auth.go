package controller

import (
	"net/http"

	"github.com/Hulhay/jk-pengker/model"
	"github.com/Hulhay/jk-pengker/shared"
	"github.com/Hulhay/jk-pengker/usecase"
	"github.com/Hulhay/jk-pengker/usecase/auth"
	"github.com/Hulhay/jk-pengker/usecase/token"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authUC  usecase.Auth
	tokenUC usecase.Token
}

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

func NewAuthController(authUC usecase.Auth, tokenUC usecase.Token) AuthController {
	return &authController{
		authUC:  authUC,
		tokenUC: tokenUC,
	}
}

func (c *authController) Register(ctx *gin.Context) {

	var (
		register auth.RegisterRequest
		err      error
	)

	err = ctx.ShouldBind(&register)
	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err = c.authUC.Register(ctx, register)
	if err != nil {
		res := shared.BuildErrorResponse("Register Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := shared.BuildResponse("Register Success!", nil)
	ctx.JSON(http.StatusCreated, res)
}

func (c *authController) Login(ctx *gin.Context) {

	var (
		response *token.ResultResponse
		login    auth.LoginRequest
		user     *model.User
		err      error
	)

	err = ctx.ShouldBind(&login)
	if err != nil {
		res := shared.BuildErrorResponse("Failed to process request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	user, err = c.authUC.Login(ctx, login)
	if err != nil {
		res := shared.BuildErrorResponse("Login Failed!", err.Error())
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response, _ = c.tokenUC.GenerateToken(ctx, user)

	res := shared.BuildResponse("Login Success!", response)
	ctx.JSON(http.StatusOK, res)
}
