package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    service "ginLearn/myweb/services"
    "strconv"
)


// @列表页面数据
// @Description get data
// @Accept  json
// @Produce json
// @Success 200 {string} string "hello"
// @Router /hello/ [get]
func GetDataList(c *gin.Context) {

    //返回结果
    c.JSON(http.StatusOK, gin.H{
        "data": "hello",
    })
}

// @增加用户
// @Description add user
// @Accept  json
// @Produce json
// @Param Id query int false "int valid"
// @Param Name query string false "string valid"
// @Param Age query int false "int valid"
// @Router /user/add [post]
func InsertUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Request.FormValue("Id"))
    name := c.Request.FormValue("Name")
    age, _ := strconv.Atoi(c.Request.FormValue("Age"))
    service.InsertUser(id, name, age)
}

// @提交LSF作业
// @Description submit lsf job
// @Accept  json
// @Produce json
// @Param jobName query string false "string valid"
// @Param cwd query string false "string valid"
// @Param cmd query string false "string valid"
// @Router /job/sub [post]
func SubmitLSFJob(c *gin.Context) {
    jobName := c.Request.FormValue("jobName")
    cwd := c.Request.FormValue("cwd")
    cmd := c.Request.FormValue("cmd")
    service.SubmitLSFJob(jobName, cwd, cmd)
}
