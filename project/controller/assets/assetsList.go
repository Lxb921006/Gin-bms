package assets

import "github.com/gin-gonic/gin"

type AssetsListForm struct {
	Page int `form:"page" validate:"min=1" binding:"required"`
}

func (a *AssetsListForm) List(ctx *gin.Context) (err error) {
	return
}
