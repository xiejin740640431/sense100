package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sense100/model/response"
	"sense100/constant"
	"sense100/model/request"
	"sense100/service"
	"sense100/model"
	"strconv"
)

//@Summary 创建承运商
//@Description 创建承运商
//@Accept json
//@Produce json
//@Param createOrUpdateCarrierReq body request.CreateOrUpdateCarrierReq true "createOrUpdateCarrierReq"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400	{object} response.BaseResponseSwag
//@Tags Carrier
//@Router /carrier/createOrUpdateCarrier [post]
func createOrUpdateCarrier(c *gin.Context) {
	var createOrUpdateCarrierReq request.CreateOrUpdateCarrierReq
	if c.Bind(&createOrUpdateCarrierReq) {
		c.JSON(http.StatusOK, service.CreateOrUpdateCarrier(&createOrUpdateCarrierReq))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 创建和更新套餐
//@Description 创建和更新套餐
//@Accept json
//@Produce json
//@Param ssSetMeal body model.SsSetMeal true "ssSetMeal"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400	{object} response.BaseResponseSwag
//@Tags Carrier
//@Router /carrier/createOrUpdateSetMeal [post]
func createOrUpdateSetMeal(c *gin.Context) {
	var setMeal model.SsSetMeal
	if c.Bind(&setMeal) {
		c.JSON(http.StatusOK, service.CreateOrUpdateSetMeal(&setMeal))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}

//@Summary 根据承运商Id和模板Id获取套餐信息
//@Description 根据承运商Id和模板Id获取套餐信息
//@Accept json
//@Produce json
//@Param carrierId path int true "carrierId"
//@Param mouldId path int true "mouldId"
//@Success 200 {array} response.BaseResponseSwag
//@Failure 400	{object} response.BaseResponseSwag
//@Tags Carrier
//@Router /carrier/getSetMealInfo/{carrierId}/{mouldId} [get]
func getSetMealInfo(c *gin.Context) {
	carrierId, err := strconv.ParseInt(c.Params.ByName("carrierId"), 10, 64)
	mouldId, err := strconv.ParseInt(c.Params.ByName("mouldId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	} else {
		c.JSON(http.StatusOK, service.FindSetMealInfo(carrierId, mouldId))
	}
}
