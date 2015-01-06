package APIs

import (
	. "github.com/julycw/ZhenhaiHRaSSB/controllers"
	"log"
	"time"
)

type DataBagStatus int

const (
	DataBagStatusSuccess DataBagStatus = iota
	DataBagStatusFailed
)

type DataBag struct {
	Status   DataBagStatus
	Data     interface{}
	Message  string
	Datetime time.Time
}

type APIController struct {
	BaseController
}

func (this *APIController) Get() {
	log.Printf("action:%v", "APIController Get")
	this.Ctx.ResponseWriter.Write([]byte("test"))
}
