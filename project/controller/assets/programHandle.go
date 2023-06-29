package assets

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/Lxb921006/Gin-bms/project/command/command"
	"github.com/Lxb921006/Gin-bms/project/dao"
	"github.com/Lxb921006/Gin-bms/project/model"
	"github.com/Lxb921006/Gin-bms/project/service"
	"github.com/Lxb921006/Gin-bms/project/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type RunProgramApiForm struct {
	Ip         string `form:"ip" json:"ip" gorm:"not null" binding:"required"`
	UpdateName string `form:"update_name" json:"update_name" gorm:"not null" binding:"required"`
	Uuid       string `form:"uuid" json:"uuid" gorm:"not null;unique" binding:"required"`
}

func (apf *RunProgramApiForm) Data() (data map[string]interface{}, err error) {
	b, err := json.Marshal(apf)
	if err != nil {
		return
	}

	if err = json.Unmarshal(b, &data); err != nil {
		return
	}

	return

}

func (apf *RunProgramApiForm) Run(ctx *gin.Context) (err error) {
	if err = ctx.ShouldBind(apf); err != nil {
		return
	}

	cy := utils.NewCelery()
	cy.Task(apf)
	close(cy.Works)

	return
}

type ProgramUpdateListForm struct {
	Ip         string `form:"ip,omitempty" json:"ip"`
	Uuid       string `form:"uuid,omitempty" json:"uuid"`
	UpdateName string `form:"update_name,omitempty" json:"update_name"`
	Project    string `form:"project,omitempty" json:"project"`
	Operator   string `form:"operator,omitempty" json:"operator"`
	Progress   int32  `form:"progress,omitempty" json:"progress"`
	Status     int32  `form:"status,omitempty" json:"status"`
	Page       int    `form:"page" json:"page" validate:"min=1" binding:"required"`
}

func (apul *ProgramUpdateListForm) List(ctx *gin.Context) (data *service.Paginate, err error) {
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

type CreateUpdateProgramRecordForm struct {
	DataList []model.AssetsProcessUpdateRecordModel `form:"data_list" json:"data_list" binding:"required"`
}

func (c *CreateUpdateProgramRecordForm) Create(ctx *gin.Context) (err error) {
	var cm model.AssetsProcessUpdateRecordModel
	if err = ctx.ShouldBindJSON(c); err != nil {
		return
	}

	if err = cm.Create(c.DataList); err != nil {
		return
	}

	return
}

type GetMissionStatusForm struct {
	Result string `form:"result" binding:"required"`
}

func (ps *GetMissionStatusForm) Get(ctx *gin.Context) (data map[string]string, err error) {
	if err = ctx.ShouldBind(ps); err != nil {
		return
	}

	data, err = dao.Rds.GetProcessStatus()
	//log.Println(data)
	if err != nil {
		return
	}

	return
}

type UploadForm struct {
	File    []string `form:"file" json:"file" binding:"required"`
	resChan chan string
}

func NewUploadForm() *UploadForm {
	return &UploadForm{
		resChan: make(chan string),
	}
}

func (u *UploadForm) UploadFiles(ctx *gin.Context) (md5 map[string]string, err error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return
	}

	files := form.File["file"]
	ips := form.Value["ips"]

	if len(files) == 0 {
		return md5, errors.New("上传失败")
	}

	var fileMd5 = make(map[string]string)

	for _, file := range files {
		fullFile := filepath.Join("C:\\Users\\Administrator\\Desktop\\update", file.Filename)
		if err = ctx.SaveUploadedFile(file, fullFile); err != nil {
			return
		}

		for _, ip := range ips {
			wg.Add(1)
			go func(ip string) {
				if err = u.SendFileToBackEnd(ip, fullFile); err != nil {
					return
				}
			}(ip)
		}
	}

	go func() {
		wg.Wait()
		close(u.resChan)
	}()

	for data := range u.resChan {
		fd := strings.Split(data, "|")
		//log.Println(fd)
		fileMd5[fd[0]] = fd[1]
	}

	md5 = fileMd5

	return
}

func (u *UploadForm) SendFileToBackEnd(ip, file string) (err error) {
	defer wg.Done()

	server := fmt.Sprintf("%s:12306", ip)

	conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	defer conn.Close()

	c := pb.NewFileTransferServiceClient(conn)

	stream, err := c.SendFile(context.Background())

	if err != nil {
		return
	}

	buffer := make([]byte, 8092)

	f, err := os.Open(file)
	if err != nil {
		return
	}

	defer f.Close()

	for {
		b, err := f.Read(buffer)
		if err == io.EOF {
			break
		}

		if b == 0 {
			break
		}

		if err = stream.Send(&pb.FileMessage{Byte: buffer[:b], Name: filepath.Base(file)}); err != nil {
			return err
		}
	}

	stream.CloseSend()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		u.resChan <- ip + "-" + resp.GetName()
	}

	return
}
