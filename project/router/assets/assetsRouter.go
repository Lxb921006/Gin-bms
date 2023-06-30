package assets

import (
	ac "github.com/Lxb921006/Gin-bms/project/controller/assets"
	"github.com/gin-gonic/gin"
)

func AssetsRouter(r *gin.Engine) {
	assets := r.Group("/assets")
	{
		assets.GET("/ws", ac.RunProgramWsController)
		assets.GET("/file/ws", ac.SyncFilePassWsController)
		assets.GET("/process/status", ac.GetMissionStatusController)
		assets.GET("/process/update/list", ac.ProgramUpdateListController)
		assets.GET("/list", ac.AssetsListController)
		assets.POST("/process/update/create", ac.CreateUpdateProgramRecordController)
		assets.POST("/api", ac.RunProgramApiController)
		assets.POST("/upload", ac.UploadController)
		assets.POST("/add", ac.AssetsCreateController)
		assets.POST("/del", ac.AssetsDeleteController)
	}
}
