package assets

import (
	"fmt"
	"github.com/Lxb921006/Gin-bms/project/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ProcessWs(ctx *gin.Context) {
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade:", err)
		return
	}

	defer conn.Close()

	ws := service.NewWs(conn)

	if err = ws.Run(); err != nil {
		if err = ws.Conn.WriteMessage(1, []byte(fmt.Sprintf("%s", err.Error()))); err != nil {
			return
		}
		return
	}
}

func ProcessApi(ctx *gin.Context) {
	var ps AssetsProcessRunForm
	if err := ps.Run(ctx); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10001,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "提交成功",
		"code":    10000,
	})

}

func GetMissionStatus(ctx *gin.Context) {
	var ps ProcessStatusForm
	data, err := ps.Get(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10001,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "ok",
		"code":    10000,
	})
}

func CreateUpdateProcess(ctx *gin.Context) {
	var create AssetsProcessRunCreateForm
	if err := create.Create(ctx); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10002,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"code":    10000,
	})
}

func UpdateListController(ctx *gin.Context) {
	var apul AssetsProcessUpdateListForm
	data, err := apul.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10001,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":     data.ModelSlice,
		"total":    data.Total,
		"pageSize": data.PageSize,
		"code":     10000,
	})
}

func AssetsListController(ctx *gin.Context) {
	var alc AssetsListForm
	data, err := alc.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10001,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":     data.ModelSlice,
		"total":    data.Total,
		"pageSize": data.PageSize,
		"code":     10000,
	})
}

func AssetsUpoadController(ctx *gin.Context) {
	var auf AssetsUpoadForm
	if err := auf.UploadFiles(ctx); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10001,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "upload ok",
		"code":    10000,
	})
}

func AssetsCreateController(ctx *gin.Context) {
	var acf AssetsCreateForm
	if err := acf.Create(ctx); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10001,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"code":    10000,
	})
}

func AssetsDeleteController(ctx *gin.Context) {
	var adf AssetsDelForm
	if err := adf.Del(ctx); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"code":    10001,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "删除成功",
		"code":    10000,
	})
}
