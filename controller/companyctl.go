package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sense100/model/response"
	"sense100/constant"
	"sense100/service"
	"strconv"
	"sense100/model"
	"log"
	"encoding/json"
)

//@Summary 添加或更新企业门面信息
//@Description 添加或更新企业门面信息
//@Accept json
//@Produce json
//@Param ssCompanyFacadeInfo body model.SsCompanyFacadeInfo true "ssCompanyFacadeInfo"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags	Company
//@Router /company/createOrUpdateFacade [post]
func createOrUpdateFacade(c *gin.Context) {
	var facadeInfo model.SsCompanyFacadeInfo
	if c.Bind(&facadeInfo) {
		bs, _ := json.Marshal(&facadeInfo)
		log.Println(string(bs))
		c.JSON(service.CreateOrUpdateFacade(&facadeInfo))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 根据用户ID获取企业门面信息
//@Description 根据用户ID获取企业门面信息
//@Accept json
//@Produce json
//@Param userId path int true "userId"
//@Success 200 {object} model.CompanyFacadeInfoResp
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Company
//@Router /company/getFacadeOnUserId/{userId} [get]
func getFacadeOnUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Params.ByName(constant.UserId))
	if err != nil {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	} else {
		c.JSON(service.FindFacadeOnUserId(userId))
	}
}

//@Summary 创建或更新行业类型
//@Description 创建或更新行业类型
//@Accept json
//@Produce json
//@Param trade body model.SsTrade true "trade"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400	{object} response.BaseResponseSwag
//@Tags Company
//@Router /company/createOrUpdateTrade [post]
func createOrUpdateTrade(c *gin.Context) {
	var trade model.SsTrade
	if c.Bind(&trade) {
		c.JSON(http.StatusOK, service.CreateOrUpdateTrade(&trade))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 根据父级Id获取行业类型集合
//@Description 根据父级Id获取行业类型集合
//@Accept json
//@Produce json
//@Param parentId path int true "parentId"
//@Success 200 {array} model.SsTrade
//@Failure 400	{object} response.BaseResponseSwag
//@Tags Company
//@Router /company/getTradesOnParentId/{parentId} [get]
func getTradesOnParentId(c *gin.Context) {
	parentId, err := strconv.ParseInt(c.Params.ByName(constant.ParentId), 10, 64)
	if err == nil {
		c.JSON(http.StatusOK, service.FindTradeOnParentId(parentId))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}
