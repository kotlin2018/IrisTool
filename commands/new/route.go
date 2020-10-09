package commands

var route = `package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"{{.Appname}}/web/controllers"
)

func InitRouter(app *iris.Application) {

	mvc.New(app.Party("/")).Handle(controllers.NewTestController())

}`
