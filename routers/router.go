package routers

import (
	"github.com/astaxie/beego"
	"github.com/julycw/ZhenhaiHRaSSB/controllers"
	"github.com/julycw/ZhenhaiHRaSSB/controllers/APIs"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// catalog
	beego.Router("/catalog", &controllers.CatalogController{})
	beego.Router("/catalog/:cid:string", &controllers.CatalogController{}, "get:GetCatalogArticle")
	// article
	beego.Router("/article/:gid:string", &controllers.ArticleController{}, "get:GetArticle")

	// APIs start!
	beego.Router("/api", &APIs.APIController{})
	//catalog
	beego.Router("/api/catalog", &APIs.CatalogAPIController{})
	//catalog - get catalog list with params
	beego.Router("/api/catalog/get_catalog_list/", &APIs.CatalogAPIController{}, "get:GetCatalogList")
	//article
	beego.Router("/api/article", &APIs.ArticleAPIController{})
	//article - - get article list with params
	beego.Router("/api/article/get_article_list/", &APIs.ArticleAPIController{}, "get:GetArticleList")
	//article - - get article one with params
	beego.Router("/api/article/get_article_one/:gid:string", &APIs.ArticleAPIController{}, "get:GetArticleOne")
	//article - - get article icon with params
	beego.Router("/api/article/get_article_icon/:gid:string/:size:string", &APIs.ArticleAPIController{}, "get:GetArticleIcon")
	// APIs end!
}
