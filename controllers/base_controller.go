package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"os"
)

const (
	PageSize  = 20
	PageCount = 10
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Layout = "layout.tpl"
}

func (this *BaseController) Get() {

}

func (this *BaseController) Post() {
}

//如果传入字符串，则原样放回，否则转换成json对象后返回
func (this *BaseController) ResponseData(data interface{}) {
	var toResponse []byte
	var err error
	switch data.(type) {
	case string:
		toResponse = []byte(data.(string))
	case []byte:
		toResponse = data.([]byte)
	default:
		toResponse, err = json.Marshal(data)
		if err != nil {
			log.Println(err.Error())
		}
	}

	this.Ctx.ResponseWriter.Write(toResponse)
}

func (this *BaseController) ResponseFile(fileType, filePath string) {
	var buf []byte = make([]byte, 255)
	var fileBody []byte = make([]byte, 0)
	if file, err := os.Open(filePath); err == nil {
		for {
			if length, _ := file.Read(buf); length == 0 {
				break
			}
			for _, v := range buf {
				fileBody = append(fileBody, v)
			}
		}
		file.Close()
	} else {
		log.Println(err.Error())
	}

	this.Ctx.Output.Header("Accept-Ranges", "bytes")
	this.Ctx.Output.Header("Content-Type", fileType)

	this.ResponseData(fileBody)
}
