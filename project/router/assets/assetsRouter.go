package assets

import (
	ac "github.com/Lxb921006/Gin-bms/project/controller/assets"
	"github.com/gin-gonic/gin"
)

func AssetsRouter(r *gin.Engine) {
	assets := r.Group("/assets")
	{
		assets.GET("/ws", ac.ProcessWsController)
		assets.GET("/process/status", ac.GetMissionStatusController)
		assets.GET("/process/update/list", ac.UpdateListController)
		assets.GET("/list", ac.AssetsListController)
		assets.POST("/process/update/create", ac.CreateUpdateProcessController)
		assets.POST("/api", ac.ProcessApiController)
		assets.POST("/upload", ac.AssetsUpoadController)
		assets.POST("/add", ac.AssetsCreateController)
		assets.POST("/del", ac.AssetsDeleteController)
	}
}
