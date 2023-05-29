package assets

import (
	ac "github.com/Lxb921006/Gin-bms/project/controller/assets"
	"github.com/gin-gonic/gin"
)

func AssetsRouter(r *gin.Engine) {
	assets := r.Group("/assets")
	{
		assets.GET("/ws", ac.ProcessWs)
		assets.GET("/process/status", ac.GetMissionStatus)
	}
}
