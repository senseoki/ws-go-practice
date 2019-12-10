package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/senseoki/ws-go-practice/iris-basic/service"
)

func main() {
	app := iris.New()

	app.Logger().AddOutput()
	app.Logger().SetLevel("debug")
	// app.Logger().SetLevel("info")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!"})
	})

	app.Get("/users/{id:uint64}", func(ctx iris.Context) {
		id := ctx.Params().GetUint64Default("id", 0)
		ctx.Application().Logger().Infof("id: %d", id)
	})

	// Path Parameters - Built-in Dependencies
	app.Get("/hero/{to:string}", hero.Handler(hello))

	hero.Register(
		&service.MyHelloService{
			Prefix: "Service: Hello",
		},
		// &service.YourHelloService{
		// 	prefix: "Service: 안녕",
		// 	name:   "your",
		// },
		&service.MyOkService{
			Prefix: "Service: OK",
		},
	)

	// Services - Static Dependencies
	app.Get("/hero/di/{to:string}", hero.Handler(helloHandler))

	//app.Run(iris.Addr(":9500"), iris.WithoutServerError(iris.ErrServerClosed))
	app.Run(iris.Addr(":9500"))
}

func hello(to string, ctx iris.Context) string {
	ctx.Application().Logger().Infof("hello: %s", to)
	return "Hello " + to
}

func helloHandler(
	to string,
	helloService service.HelloService,
	okService service.OkService,
	ctx iris.Context) string {
	//
	ctx.Application().Logger().Infof("HelloService: %s", to)
	hello := helloService.SayHello(to)
	ok := okService.SayOK(to)
	return hello + " " + ok
}
