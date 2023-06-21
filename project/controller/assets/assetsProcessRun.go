package assets

import (
	"encoding/json"
	"errors"
	"github.com/Lxb921006/Gin-bms/project/model"
	"github.com/Lxb921006/Gin-bms/project/service"
	"github.com/Lxb921006/Gin-bms/project/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"path/filepath"
)

type AssetsProcessRunForm struct {
	Ip         string `form:"ip" json:"ip" gorm:"not null" binding:"required"`
	UpdateName string `form:"update_name" json:"update_name" gorm:"not null" binding:"required"`
	Uuid       string `form:"uuid" json:"uuid" gorm:"not null;unique" binding:"required"`
}

func (apf *AssetsProcessRunForm) Data() (data map[string]interface{}, err error) {
	b, err := json.Marshal(apf)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &data); err != nil {
		return
	}

	return

}

func (apf *AssetsProcessRunForm) Run(ctx *gin.Context) (err error) {
	if err = ctx.ShouldBind(apf); err != nil {
		return
	}

	cy := utils.NewCelery()
	cy.Task(apf)
	close(cy.Works)

	return
}

func (apf *AssetsProcessRunForm) SendFile() {}

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

type AssetsProcessRunCreateForm struct {
	DataList []model.AssetsProcessUpdateRecordModel `form:"data_list" json:"data_list" binding:"required"`
}

func (c *AssetsProcessRunCreateForm) Create(ctx *gin.Context) (err error) {
	var cm model.AssetsProcessUpdateRecordModel
	if err = ctx.ShouldBindJSON(c); err != nil {
		return
	}

	if err = cm.Create(c.DataList); err != nil {
		return
	}

	return
}

type AssetsUpoadForm struct {
	Files []string `form:"upload" json:"upload" binding:"required"`
}

func (u *AssetsUpoadForm) UploadFiles(ctx *gin.Context) (err error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}

	files := form.File["file"]
	data := form.Value["ips"]

	log.Println(data)

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

func (u *AssetsUpoadForm) SendFileToBackEnd(ip, file string) (err error) {

	return
}
