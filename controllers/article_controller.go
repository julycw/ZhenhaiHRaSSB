package controllers

import (
	"github.com/julycw/ZhenhaiHRaSSB/models"
	"github.com/julycw/orm"
	"github.com/julycw/paginator"
)

type ArticleController struct {
	BaseController
}

func (this *ArticleController) Get() {
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

	totalCount := models.ArticleStore.GetCountByCondition(conditions)
	articleList := models.ArticleStore.GetByPageAndConditionAndOrder(int(page), int(size), conditions, orders)

	pagination := paginator.New(totalCount, int(size), PageCount)

	this.Data["articleList"] = &articleList
	this.Data["pagination"] = &pagination
}

func (this *ArticleController) GetArticle() {
	this.TplNames = "article/index.tpl"
	conditions := []orm.Condition{}
	orders := []orm.Order{}

	gid := this.GetString(":gid")
	conditions = append(conditions, orm.Condition{Name: "gid", Compare: "=", Value: gid})

	articleList := models.ArticleStore.GetByPageAndConditionAndOrder(1, 1, conditions, orders)
	if len(articleList) > 0 {
		this.Data["article"] = articleList[0].(models.Article)
	}
}
