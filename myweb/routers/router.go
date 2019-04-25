package routers

import (
  "github.com/gin-gonic/gin"
  . "ginLearn/myweb/controllers" //api部分
 )

func InitRouter() *gin.Engine{
    router := gin.Default()
    //Hello World
    router.GET("/hello", GetDataList)

    router.POST("/user/add", InsertUser)

    router.POST("/job/sub", SubmitLSFJob)

    return router
}
