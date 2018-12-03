package controller

import (
	"github.com/gin-gonic/gin"
	"sense100/model"
	"net/http"
	"sense100/service"
	"sense100/model/response"
	"sense100/constant"
	"sense100/model/request"
)

//@Summary 发布动态
//@Description 发布动态
//@Accept json
//@Produce json
//@Param ssSpDynamic body model.SsSpDynamic true "ssSpDynamic"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400	{object} response.BaseResponseSwag
//@Tags Company
//@Router /company/publishDynamic [post]
func publishDynamic(c *gin.Context) {
	var dynamic model.SsSpDynamic
	if c.Bind(&dynamic) {
		c.JSON(http.StatusOK, service.CreateDynamic(&dynamic))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 分页获取动态
//@Description 分页获取动态
//@Accept json
//@Produce json
//@Param getDynamicListPageReq body request.GetDynamicListPageReq true "getDynamicListPageReq"
//@Success 200 {array} model.DynamicResp
//@Failure 400	{object} response.BaseResponseSwag
//@Tags Company
//@Router /company/getDynamicListPage [post]
func getDynamicListPage(c *gin.Context) {
	var req request.GetDynamicListPageReq
	if c.Bind(&req) {
		c.JSON(http.StatusOK,service.FindDynamicsOnPage(&req))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}
