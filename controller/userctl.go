package controller

import (
	"github.com/gin-gonic/gin"
	"sense100/model/request"
	"net/http"
	"sense100/model/response"
	"sense100/constant"
	"sense100/service"
	"sense100/model"
)

//@Summary 注册用户
//@Description 注册用户 注册类型必传，手机号注册：手机号、密码必传 微信注册：微信id必传
//@Accept json
//@Produce json
//@Param registerUserReq body request.RegisterUserReq	 true "register"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags User
//@Router	/user/registerUser	[post]
func registerUser(c *gin.Context) {
	var registerUserReq request.RegisterUserReq
	if !c.Bind(&registerUserReq) {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	} else {
		c.JSON(service.RegisterUser(&registerUserReq))
	}
}

//@Summary 登录
//@Description 登录 登录类型必传 手机号登录：手机号、密码必传  微信登录：微信id必传
//@Accept json
//@Produce json
//@Param loginReq body request.LoginReq	 true "loginReq"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags User
//@Router	/user/login	[post]
func login(c *gin.Context) {
	var loginReq request.LoginReq
	if c.Bind(&loginReq) {
		c.JSON(http.StatusOK, service.FindUser(&loginReq))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 添加浏览记录
//@Description 添加浏览记录
//@Accept json
//@Produce json
//@Param addBrowseRecordReq body model.AddBrowseRecordReq true "addBrowseRecordReq"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags User
//@Router /user/addBrowseRecord [post]
func addBrowseRecord(c *gin.Context) {
	var req model.AddBrowseRecordReq
	if c.Bind(&req) {
		c.JSON(http.StatusOK, service.CreateOrUpdateBrowseRecord(&req))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 添加点赞记录
//@Description 添加点赞记录
//@Accept json
//@Produce json
//@Param praise body model.SsSpPraise true "praise"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags User
//@Router /user/doPraise [post]
func doPraise(c *gin.Context) {
	var praise model.SsSpPraise
	if c.Bind(&praise) {
		c.JSON(http.StatusOK, service.CreatePraise(&praise))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}
