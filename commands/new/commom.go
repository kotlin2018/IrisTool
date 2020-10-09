package commands

var common = `
package controllers

import "github.com/kataras/iris/v12"

//持有一个iris上下文实例
type Common struct {
	Ctx iris.Context
}

// 用于前后端交互数据的json
type JsonStruct struct {
	Code  int         //json:"code"
	Msg   interface{} //json:"msg"
	Items interface{} //json:"items"
	Count int64       //json:"count"
}

func (this *Common) RetrunError(code int, msg interface{}) {
	json := &JsonStruct{
		Code: code,
		Msg: msg,
	}
	this.Ctx.JSON(json)
	this.Ctx.StopExecution()
	return
}

func (this *Common) ReturnSuccess(code int,msg interface{},items interface{},count int64) {
	json := &JsonStruct{
		Code: code,
		Msg: msg,
		Items: items,
		Count: count,
	}
	this.Ctx.JSON(json)
	this.Ctx.StopExecution()
	return
}
`