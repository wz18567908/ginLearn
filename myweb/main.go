package main

import (
	_ "ginLearn/myweb/docs"
    router "ginLearn/myweb/routers"
    db "ginLearn/myweb/db"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

// @host 192.168.0.5:8080
func main() {
    db.InitDB()
    router:=router.InitRouter()

    // use ginSwagger middleware to 
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    router.Run()
}
