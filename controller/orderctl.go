package controller

import (
	"github.com/gin-gonic/gin"
	"sense100/model/request"
	"net/http"
	"sense100/model/response"
	"sense100/constant"
	"sense100/service"
)

//@Summary 提交订单
//@Description 提交订单
//@Accept json
//@Produce json
//@Param submitOrderReq body request.SubmitOrderReq true "submitOrderReq"
//@Success 200 {object} response.BaseResponseSwag
//@Failure 400 {object} response.BaseResponseSwag
//@Tags Order
//@Router /order/submitOrder [post]
func submitOrder(c *gin.Context) {
	var submitOrderReq request.SubmitOrderReq
	if c.Bind(&submitOrderReq) {
		c.JSON(http.StatusOK, service.CreateOrderInfo(&submitOrderReq))
	} else {
		c.JSON(http.StatusOK, response.RespWithStatus(constant.ErrParamDefect))
	}
}
