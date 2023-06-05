package assets

import (
	"errors"
	"fmt"
	"github.com/Lxb921006/Gin-bms/project/command/client"
	"github.com/Lxb921006/Gin-bms/project/model"
	"github.com/Lxb921006/Gin-bms/project/service"
	"github.com/Lxb921006/Gin-bms/project/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"path/filepath"
	"time"
)

// 远程调用对应脚本
type AssetsProcessRunForm struct {
	Ip         string     `form:"ip" json:"ip" gorm:"not null" binding:"required"`
	UpdateName string     `form:"update_name" json:"update_name" gorm:"not null" binding:"required"`
	Uuid       string     `form:"uuid" json:"uuid" gorm:"not null;unique" binding:"required"`
	Err        chan error `form:"err,omitempty" json:"-"`
}

func NewAssetsProcessRunForm() *AssetsProcessRunForm {
	return &AssetsProcessRunForm{
		Err: make(chan error, 1),
	}
}

func (apf *AssetsProcessRunForm) Run(ctx *gin.Context) (err error) {
	if err = ctx.ShouldBind(apf); err != nil {
		return
	}

	go func() {
		server := fmt.Sprintf("%s:12306", apf.Ip)
		conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			apf.Err <- err
			return
		}

		cn := client.NewRpcClient(apf.UpdateName, apf.Uuid, nil, conn)
		go func() {
			if err = cn.Send(); err != nil {
				apf.Err <- err
			}
		}()
	}()

	for {
		select {
		case err = <-apf.Err:
			return
		case <-time.After(time.Second):
			return
		}
	}
}

// 更新列表查询
type AssetsProcessUpdateListForm struct {
	Ip         string `form:"ip,omitempty" json:"ip"`
	Uuid       string `form:"uuid,omitempty" json:"uuid"`
	UpdateName string `form:"update_name,omitempty" json:"update_name"`
	Project    string `form:"project,omitempty" json:"project"`
	Operator   string `form:"operator,omitempty" json:"operator"`
	Progress   int32  `form:"progress,omitempty" json:"progress"`
	Status     int32  `form:"status,omitempty" json:"status"`
	Page       int    `form:"page" json:"page" validate:"min=1" binding:"required"`
}

func (apul *AssetsProcessUpdateListForm) List(ctx *gin.Context) (data *service.Paginate, err error) {
	var lm model.AssetsProcessUpdateRecordModel
	if err = ctx.ShouldBind(apul); err != nil {
		return
	}

	validate := validator.New()
	vd := NewValidateData(validate)
	if err = vd.ValidateStruct(apul); err != nil {
		return
	}

	if err = utils.CopyStruct(apul, &lm); err != nil {
		return
	}

	data, err = lm.List(apul.Page, lm)
	if err != nil {
		return
	}

	return
}

// 更新列表添加数据
type AssetsProcessRunCreateForm struct {
}

func (c *AssetsProcessRunCreateForm) Create(ctx *gin.Context) (err error) {
	var cm model.AssetsProcessUpdateRecordModel
	if err = ctx.ShouldBindJSON(&cm); err != nil {
		return
	}

	if err = cm.Create(cm); err != nil {
		return
	}

	return
}

// 上传文件
type AssetsUpoadForm struct {
	Upload []string `form:"upload" json:"upload" binding:"required"`
}

func (u *AssetsUpoadForm) UploadFiles(ctx *gin.Context) (err error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}

	files := form.File["file"]
	if len(files) == 0 {
		return errors.New("上传失败")
	}

	for _, file := range files {
		if err = ctx.SaveUploadedFile(file, filepath.Join("C:\\Users\\Administrator\\Desktop\\update", file.Filename)); err != nil {
			return
		}
	}

	return
}
