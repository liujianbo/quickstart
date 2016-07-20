package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}


beego.Get("/",func(ctx *context.Context){
     ctx.Output.Body([]byte("hello world"))
})


beego.Post("/alice",func(ctx *context.Context){
     ctx.Output.Body([]byte("bob"))
})

beego.Any("/foo",func(ctx *context.Context){
     ctx.Output.Body([]byte("bar"))
})

s := rpc.NewServer()
s.RegisterCodec(json.NewCodec(), "application/json")
s.RegisterService(new(HelloService), "")
beego.Handler("/rpc", s)