package api

import "github.com/gin-gonic/gin"

type Api interface {
	ValidateStruct(any interface{}) (err error)
}

type FillDataInterface interface {
	FillData() (any interface{}, err error)
}

type CeleryInterface interface {
	Run(ctx *gin.Context) (err error)
}
