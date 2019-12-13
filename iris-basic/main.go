package main 

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/senseoki/ws-go-practice/iris-basic/service"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	_ "github.com/senseoki/ws-go-practice/iris-basic/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	app := newApp()

	//app.Run(iris.Addr(":9500"), iris.WithoutServerError(iris.ErrServerClosed))
	//app.Run(iris.TLS(":443", "mycert.cert", "mykey.key"))
	app.Run(iris.Addr(":9500"))
	
}

func newApp() *iris.Application {
	app := iris.New()

	app.Logger().AddOutput()
	app.Logger().SetLevel("debug")
	// app.Logger().SetLevel("info")

	app.Use(recover.New())
	app.Use(logger.New())

	swaggerConfig := &swagger.Config{
        URL: "http://localhost:9500/swagger/doc.json", //The url pointing to API definition
    }
    // use swagger middleware to 
    app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler))

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
	app.Get("/hero/{to:string}", hero.Handler(Hello))

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
	app.Get("/hero/di/{to:string}", hero.Handler(TestService))

	return app
}

// Hello ...
func Hello(to string, ctx iris.Context) string {
	ctx.Application().Logger().Infof("hello: %s", to)
	return "Hello " + to
}

// TestService ...
func TestService( 
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
