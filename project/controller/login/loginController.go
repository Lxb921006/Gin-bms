package login

import (
	"fmt"
	"net/http"

	"github.com/Lxb921006/Gin-bms/project/model"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	UserName string `form:"user" binding:"required" json:"user"`
	Password string `form:"password" binding:"required"`
}

type GaloginForm struct {
	UserName string `form:"user" binding:"required" json:"user"`
	Code     string `form:"code" binding:"required"`
}

func Galogin(ctx *gin.Context) {
	var ga GaloginForm
	var l model.Login

	if err := ctx.ShouldBind(&ga); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := l.GaLogin(ga.Code, ga.UserName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": fmt.Sprintf("登录成功, 欢迎%s大佬!!!", ga.UserName),
	})

}

func Login(ctx *gin.Context) {
	var l model.Login
	var lf LoginForm

	if err := ctx.ShouldBind(&lf); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := l.UserLogin(lf.UserName, lf.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := lf.UserName
	ip := ctx.RemoteIP()
	url := ctx.Request.URL.Path
	op := model.OperateLog{}
	if err := op.AddOperateLog(url, user, ip); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": fmt.Sprintf("欢迎%s大佬!!!", lf.UserName),
	})

}

func Logout(ctx *gin.Context) {
	var l model.Login
	user := ctx.PostForm("user")
	if user == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "请选择用户退出",
		})
		return
	}

	if err := l.UserLogout(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ip := ctx.RemoteIP()
	url := ctx.Request.URL.Path
	op := model.OperateLog{}
	if err := op.AddOperateLog(url, user, ip); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s退出成功, 欢迎再次光临!", user),
	})

}
