package main

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"

	_ "lab150/lab004/lab002/docs"
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
	app := iris.New()

	url := swagger.URL("http://localhost:8080/swagger/doc.json") //The url pointing to API definition
	app.Get("/swagger/{any:path}", swagger.WrapHandler(swaggerFiles.Handler, url))

	app.Run(iris.Addr(":8080"))
}
