package main

import (
	"github.com/astaxie/beego"
	_ "github.com/julycw/ZhenhaiHRaSSB/routers"
	// "log"
)

func main() {
	beego.Run()
}

func init() {
	// log.SetFlags(log.Lshortfile | log.LstdFlags)
}
