package controller

import (
	"net/http"

	"github.com/Hulhay/jk-pengker/shared"
	"github.com/Hulhay/jk-pengker/usecase"
	"github.com/gin-gonic/gin"
)

type storeController struct {
	storeUC usecase.Store
	tokenUC usecase.Token
}

type StoreController interface {
	GetStoreList(ctx *gin.Context)
}

func NewStoreController(storeUc usecase.Store, tokenUC usecase.Token) StoreController {
	return &storeController{
		storeUC: storeUc,
		tokenUC: tokenUC,
	}
}

func (c *storeController) GetStoreList(ctx *gin.Context) {

	store, err := c.storeUC.GetStoreList(ctx)
	if err != nil {
		res := shared.BuildErrorResponse("Get Store Failed!", err.Error())
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	res := shared.BuildResponse("Success", store)
	ctx.JSON(http.StatusOK, res)
}
