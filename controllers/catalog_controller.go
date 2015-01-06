package controllers

import (
	"github.com/julycw/ZhenhaiHRaSSB/models"
	"github.com/julycw/orm"
	"github.com/julycw/paginator"
)

type CatalogController struct {
	BaseController
}

func (this *CatalogController) Prepare() {
	this.BaseController.Prepare()
	this.TplNames = "catalog/index.tpl"
}

func (this *CatalogController) Get() {
	this.TplNames = "catalog/index.tpl"

	conditions := []orm.Condition{}
	orders := []orm.Order{}

	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	size, _ := this.GetInt("size")
	if size == 0 {
		size = PageSize
	}

	totalCount := models.CatalogStore.GetCountByCondition(conditions)
	catalogList := models.CatalogStore.GetByPageAndConditionAndOrder(int(page), int(size), conditions, orders)

	pagination := paginator.New(totalCount, int(size), PageCount)

	this.Data["catalogList"] = &catalogList
	this.Data["pagination"] = &pagination
}

func (this *CatalogController) GetCatalogArticle() {
	this.TplNames = "catalog/catalog_articles.tpl"

	conditions := []orm.Condition{}
	orders := []orm.Order{}

	cid := this.GetString(":cid")
	if cid != "" {
		conditions = append(conditions, orm.Condition{Name: "cid", Compare: "=", Value: cid})
	}

	page, _ := this.GetInt("page")
	if page == 0 {
		page = 1
	}
	size, _ := this.GetInt("size")
	if size == 0 {
		size = PageSize
	}

	totalCount := models.ArticleStore.GetCountByCondition(conditions)
	articleList := models.ArticleStore.GetByPageAndConditionAndOrder(int(page), int(size), conditions, orders)

	pagination := paginator.New(totalCount, int(size), PageCount)

	this.Data["articleList"] = &articleList
	this.Data["pagination"] = &pagination
}
