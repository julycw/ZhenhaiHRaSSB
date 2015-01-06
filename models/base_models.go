package models

import (
	"github.com/astaxie/beego"
	"github.com/julycw/orm"
)

var (
	ArticleStore         orm.DBStore
	CatalogStore         orm.DBStore
	ArticleResourceStore orm.DBStore
)

func init() {
	dbName := beego.AppConfig.String("dbname")
	dbUser := beego.AppConfig.String("dbuser")
	dbPasswd := beego.AppConfig.String("dbPasswd")
	dbHost := beego.AppConfig.String("dbHost")

	orm.RegisterOrm(dbHost, dbUser, dbPasswd, dbName)

	ArticleStore, _ = orm.GetDBStore("Cell_Article", &Article{})
	CatalogStore, _ = orm.GetDBStore("Cell_Catalog", &Catalog{})
	ArticleResourceStore, _ = orm.GetDBStore("Cell_Resource", &ArticleResource{})
}
