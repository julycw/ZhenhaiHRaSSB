package APIs

import (
	. "github.com/julycw/ZhenhaiHRaSSB/controllers"
	"github.com/julycw/ZhenhaiHRaSSB/models"
	"github.com/julycw/orm"
	"log"
	"strings"
	"time"
)

type CatalogAPIController struct {
	BaseController
}

func (this *CatalogAPIController) Get() {
	log.Printf("action:%v", "CatalogAPIController Get")
	this.ResponseData("data")
}

func (this *CatalogAPIController) GetCatalogList() {
	orders := []orm.Order{}
	pageIndex, _ := this.GetInt("page")
	if pageIndex <= 0 {
		pageIndex = 1
	}
	pageSize, _ := this.GetInt("size")
	if pageSize <= 0 {
		pageSize = PageSize
	}

	orderby := this.GetString("orderby")
	if orderby != "" {
		if strings.HasPrefix(orderby, "-") {
			orders = append(orders, orm.Order{Name: orderby[1:], By: orm.DESC})
		} else {
			orders = append(orders, orm.Order{Name: orderby, By: orm.ASC})
		}
	}

	catalogList := models.CatalogStore.GetByPageAndConditionAndOrder(int(pageIndex), int(pageSize), []orm.Condition{}, orders)

	dataBag := DataBag{
		Message:  "",
		Datetime: time.Now(),
	}
	if len(catalogList) > 0 {
		dataBag.Status = DataBagStatusSuccess
		dataBag.Data = catalogList
	} else {
		dataBag.Status = DataBagStatusFailed
	}
	this.Data["json"] = &dataBag
	this.ServeJson()
}
